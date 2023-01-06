package main

import (
	"context"
	"github.com/mholt/archiver/v4"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup tools",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute root command", err)
		os.Exit(1)
	}
}

func download(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", err)
		}
	}()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			slog.Warn("failed to close file", err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
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
