package main

import (
	"context"
	"github.com/fhluo/json2go/pkg/downloaders"
	"github.com/mholt/archiver/v4"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup tools",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			slog.Error("failed to get current directory", "err", err)
			os.Exit(1)
		}

		dir := wd
		for filepath.Base(dir) != "json2go" && filepath.Dir(dir) != dir {
			dir = filepath.Dir(dir)
		}

		if filepath.Dir(dir) == dir {
			slog.Error("failed to find json2go directory", "working directory", wd)
			os.Exit(1)
		}

		installPath := filepath.Join(dir, "build", "tools")

		if upx {
			installUPX("", githubToken, installPath)
		}
		if nsis {
			installNSIS("", installPath)
		}
	},
}

var (
	upx         bool
	nsis        bool
	githubToken string
)

func init() {
	rootCmd.Flags().BoolVar(&upx, "upx", false, "install upx")
	rootCmd.Flags().BoolVar(&nsis, "nsis", false, "install nsis")

	rootCmd.PersistentFlags().StringVarP(&githubToken, "token", "t", "", "github token")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute root command", "err", err)
		os.Exit(1)
	}
}

func extract(ctx context.Context, filename string, dst string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			slog.Warn("failed to close file", err)
		}
	}()

	return archiver.Zip{}.Extract(ctx, r, nil, func(ctx context.Context, f archiver.File) error {
		if f.IsDir() {
			return nil
		}

		path := filepath.Join(dst, f.NameInArchive)
		err := os.MkdirAll(filepath.Dir(path), 0660)
		if err != nil {
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			return err
		}
		defer func() {
			if err := out.Close(); err != nil {
				slog.Warn("failed to close file", err)
			}
		}()

		r, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(out, r)
		return err
	})
}

func install(downloader downloaders.Downloader, installPath string) {
	packagePath, err := downloader.DownloadToTempDir()
	if err != nil {
		slog.Error("failed to download NSIS", "err", err, "url", downloader.DownloadURL(), "packagePath", packagePath)
		os.Exit(1)
	}

	if err := extract(context.Background(), packagePath, installPath); err != nil {
		slog.Error("failed to extract files", "err", err, "packagePath", packagePath, "installPath", installPath)
		os.Exit(1)
	}
}
