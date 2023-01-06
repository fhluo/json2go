package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
)

var upxCmd = &cobra.Command{
	Use:   "upx",
	Short: "setup upx",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		url, name, err := getUPXDownloadURLAndFilename(ctx, githubToken)
		if err != nil {
			slog.Error("failed to get download url", err)
			os.Exit(1)
		}

		switch _, err = os.Stat(name); {
		case os.IsNotExist(err):
			if err = download(url, name); err != nil {
				slog.Error("failed to download upx", err)
				os.Exit(1)
			}
		case err == nil:
			slog.Info("file already exists, skip download")
		default:
			slog.Error("failed to stat file", err)
			os.Exit(1)
		}

		if err = extract(ctx, name, upxPath); err != nil {
			slog.Error("failed to extract files", err, "name", name)
			os.Exit(1)
		}
	},
}

var (
	githubToken string
	upxPath     string
	upxVersion  = "4.0.1"
)

func init() {
	rootCmd.AddCommand(upxCmd)

	upxCmd.Flags().StringVarP(&githubToken, "token", "t", "", "github token")
	upxCmd.Flags().StringVarP(&upxPath, "path", "p", ".", "path to extract upx")
}

func getUPXDownloadURL() string {
	return fmt.Sprintf("https://github.com/upx/upx/releases/download/v%[1]s/upx-%[1]s-win64.zip", upxVersion)
}

func getUPXDownloadFilename() string {
	return fmt.Sprintf("upx-%s-win64.zip", upxVersion)
}

func getUPXDownloadURLAndFilename(ctx context.Context, token string) (string, string, error) {
	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(tc)

	release, _, err := client.Repositories.GetLatestRelease(ctx, "upx", "upx")
	if err != nil {
		if _, ok := err.(*github.RateLimitError); ok {
			slog.Info("hit rate limit, use default", "upx-version", upxVersion)
			return getUPXDownloadURL(), getUPXDownloadFilename(), nil
		} else {
			slog.Warn("failed to get latest release, use default", err, "upx-version", upxVersion)
			return getUPXDownloadURL(), getUPXDownloadFilename(), nil
		}
	}

	result := lo.Filter(release.Assets, func(asset *github.ReleaseAsset, _ int) bool {
		return strings.Contains(asset.GetName(), "win64")
	})
	if len(result) == 0 {
		return "", "", fmt.Errorf("failed to find win64 release")
	}

	return result[0].GetBrowserDownloadURL(), result[0].GetName(), nil
}
