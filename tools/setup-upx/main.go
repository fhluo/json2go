package main

import (
	"context"
	"fmt"
	"github.com/fhluo/json2go/pkg/util"
	"github.com/google/go-github/v49/github"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
}

var rootCmd = &cobra.Command{
	Use:  "setup-upx",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		downloadURL, filename, err := getDownloadURLAndFilename(token)
		if err != nil {
			slog.Error("failed to get download url", err)
			os.Exit(1)
		}

		switch _, err = os.Stat(filename); {
		case os.IsNotExist(err):
			if err = util.DownloadFile(downloadURL, filename); err != nil {
				slog.Error("failed to download upx", err)
				os.Exit(1)
			}
		case err == nil:
			slog.Info("file already exists, skip download")
		default:
			slog.Error("failed to stat file", err)
			os.Exit(1)
		}

		if path != "" {
			// make dir
			if err = os.MkdirAll(path, 0660); err != nil {
				slog.Error("failed to create dir", err)
				os.Exit(1)
			}

			if err = util.ExtractOne(filename, "upx.exe", filepath.Join(path, "upx.exe")); err != nil {
				slog.Error("failed to extract files", err, "filename", filename)
				os.Exit(1)
			}
		} else {
			if err = util.ExtractOne(filename, "upx.exe", "upx.exe"); err != nil {
				slog.Error("failed to extract files", err, "filename", filename)
				os.Exit(1)
			}
		}
	},
}

var (
	token string
	path  string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "github token")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to extract")
}

func getDownloadURLAndFilename(githubToken string) (downloadURL string, filename string, err error) {
	ctx := context.Background()

	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(tc)

	release, _, err := client.Repositories.GetLatestRelease(ctx, "upx", "upx")
	if err != nil {
		return "", "", fmt.Errorf("failed to get latest release: %w", err)
	}

	result := lo.Filter(release.Assets, func(asset *github.ReleaseAsset, _ int) bool {
		return strings.Contains(asset.GetName(), "win64")
	})
	if len(result) == 0 {
		return "", "", fmt.Errorf("failed to find win64 release")
	}

	return result[0].GetBrowserDownloadURL(), result[0].GetName(), nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute command", err)
		os.Exit(1)
	}
}
