package main

import (
	"context"
	"fmt"
	"github.com/mholt/archiver/v4"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const Version = "3.08"

var (
	URL      = fmt.Sprintf("https://sourceforge.net/projects/nsis/files/NSIS%%203/%s/nsis-%s.zip/download", Version, Version)
	Filename = fmt.Sprintf("nsis-%s.zip", Version)
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
}

func DownloadNSIS() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", err)
		}
	}()

	out, err := os.Create(Filename)
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

func HandleFile(ctx context.Context, f archiver.File) error {
	err := os.MkdirAll(filepath.Dir(f.NameInArchive), 0660)
	if err != nil {
		return err
	}

	out, err := os.Create(f.NameInArchive)
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
}

func Unzip(filename string) error {
	// open file
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			slog.Warn("failed to close file", err)
		}
	}()

	return archiver.Zip{}.Extract(context.Background(), r, nil, HandleFile)
}

func main() {
	// if file not exits
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		if err = DownloadNSIS(); err != nil {
			slog.Error("failed to download NSIS", err)
			os.Exit(1)
		}
	} else {
		if err != nil {
			slog.Error("failed to stat file", err)
			os.Exit(1)
		} else {
			slog.Info("file already exists, skip download")
		}
	}

	if err := Unzip(Filename); err != nil {
		slog.Error("failed to unzip NSIS", err)
		os.Exit(1)
	}
}
