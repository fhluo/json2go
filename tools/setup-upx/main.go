package main

import (
	"context"
	"github.com/fhluo/json2go/pkg/util"
	"github.com/google/go-github/v49/github"
	"github.com/samber/lo"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
}

func main() {
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "upx", "upx")
	if err != nil {
		slog.Error("failed to get latest release", err)
		os.Exit(1)
	}

	result := lo.Filter(release.Assets, func(asset *github.ReleaseAsset, _ int) bool {
		return strings.Contains(asset.GetName(), "win64")
	})
	if len(result) == 0 {
		slog.Error("failed to find win64 release", nil, "assets", lo.Map(release.Assets, func(asset *github.ReleaseAsset, _ int) string {
			return asset.GetName()
		}))
		os.Exit(1)
	}

	name, url := result[0].GetName(), result[0].GetBrowserDownloadURL()

	switch _, err = os.Stat(name); {
	case os.IsNotExist(err):
		if err = util.DownloadFile(url, name); err != nil {
			slog.Error("failed to download upx", err)
			os.Exit(1)
		}
	case err == nil:
		slog.Info("file already exists, skip download")
	default:
		slog.Error("failed to stat file", err)
		os.Exit(1)
	}

	if err = util.ExtractOne(name, "upx.exe", "upx.exe"); err != nil {
		slog.Error("failed to extract files", err, "filename", name)
		os.Exit(1)
	}
}
