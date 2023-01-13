package downloaders

import (
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Downloader interface {
	DownloadURL() string
	DownloadFilename() string
	Download(path string) error
	DownloadToTempDir() (path string, err error)
}

func Download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", err)
		}
	}()

	out, err := os.Create(path)
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

func FileExists(path string) (bool, error) {
	switch _, err := os.Stat(path); {
	case err == nil:
		return true, nil
	case os.IsNotExist(err):
		return false, nil
	default:
		return false, err
	}
}

// DownloadToTempDir downloads the file to os.TempDir() if the file does not exist.
func DownloadToTempDir(url string, filename string) (path string, err error) {
	path = filepath.Join(os.TempDir(), filename)

	exists, err := FileExists(path)
	if err != nil {
		return path, err
	}
	if exists {
		return path, nil
	}

	err = Download(url, path)
	return
}
