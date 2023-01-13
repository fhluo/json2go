package downloaders

import (
	"fmt"
)

const DefaultNSISVersion = "3.08"

type NSISDownloader struct {
	version string
}

func NewNSISDownloader(version string) *NSISDownloader {
	if version == "" {
		version = DefaultNSISVersion
	}

	return &NSISDownloader{version: version}
}

func (n *NSISDownloader) DownloadURL() string {
	return fmt.Sprintf("https://sourceforge.net/projects/nsis/files/NSIS%%203/%[1]s/nsis-%[1]s.zip/download", n.version)
}

func (n *NSISDownloader) DownloadFilename() string {
	return fmt.Sprintf("nsis-%s.zip", n.version)
}

func (n *NSISDownloader) Download(path string) error {
	return Download(n.DownloadURL(), path)
}

func (n *NSISDownloader) DownloadToTempDir() (path string, err error) {
	return DownloadToTempDir(n.DownloadURL(), n.DownloadFilename())
}
