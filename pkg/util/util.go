package util

import (
	"context"
	"github.com/mholt/archiver/v4"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, filename string) error {
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

func handleFile(ctx context.Context, f archiver.File) error {
	if f.IsDir() {
		return nil
	}

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

func Extract(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			slog.Warn("failed to close file", err)
		}
	}()

	return archiver.Zip{}.Extract(context.Background(), r, nil, handleFile)
}

func ExtractOne(filename string, target string, dst string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			slog.Warn("failed to close file", err)
		}
	}()

	return archiver.Zip{}.Extract(context.Background(), r, nil, func(ctx context.Context, f archiver.File) error {
		if f.IsDir() {
			return nil
		}

		if filepath.Base(f.NameInArchive) != target {
			return nil
		}

		out, err := os.Create(dst)
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
