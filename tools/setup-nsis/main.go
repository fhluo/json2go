package main

import (
	"fmt"
	"github.com/fhluo/json2go/pkg/util"
	"golang.org/x/exp/slog"
	"os"
)

const Version = "3.08"

var (
	URL      = fmt.Sprintf("https://sourceforge.net/projects/nsis/files/NSIS%%203/%s/nsis-%s.zip/download", Version, Version)
	Filename = fmt.Sprintf("nsis-%s.zip", Version)
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
}

func main() {
	switch _, err := os.Stat(Filename); {
	case os.IsNotExist(err):
		if err = util.DownloadFile(URL, Filename); err != nil {
			slog.Error("failed to download NSIS", err)
			os.Exit(1)
		}
	case err == nil:
		slog.Info("file already exists, skip download")
	default:
		slog.Error("failed to stat file", err)
		os.Exit(1)
	}

	if err := util.Extract(Filename); err != nil {
		slog.Error("failed to extract files", err, "filename", Filename)
		os.Exit(1)
	}
}
