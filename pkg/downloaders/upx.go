package downloaders

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

const DefaultUPXVersion = "4.0.2"

type UPXDownloader struct {
	version     string
	githubToken string

	release *github.RepositoryRelease
	apiErr  error
}

func NewUPXDownloader(version string, githubToken string) *UPXDownloader {
	if version == "" {
		version = DefaultUPXVersion
	}

	return &UPXDownloader{
		version:     version,
		githubToken: githubToken,
	}
}

func (u *UPXDownloader) GetLatestRelease() (*github.RepositoryRelease, error) {
	if u.release != nil {
		return u.release, nil
	}

	ctx := context.Background()

	var tc *http.Client
	if u.githubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: u.githubToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	client := github.NewClient(tc)

	release, _, err := client.Repositories.GetLatestRelease(ctx, "upx", "upx")
	if err != nil {
		u.apiErr = err
		return nil, err
	}

	u.release = release
	return u.release, nil
}

func (u *UPXDownloader) GetLatestReleaseAsset() (*github.ReleaseAsset, error) {
	if u.apiErr != nil {
		return nil, u.apiErr
	}

	release, err := u.GetLatestRelease()
	if err != nil {
		return nil, err
	}

	result := lo.Filter(release.Assets, func(asset *github.ReleaseAsset, _ int) bool {
		return strings.Contains(asset.GetName(), "win64")
	})
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to find a suitable release asset")
	}

	return result[0], nil
}

func (u *UPXDownloader) downloadURL() string {
	return fmt.Sprintf("https://github.com/upx/upx/releases/download/v%[1]s/upx-%[1]s-win64.zip", u.version)
}

func (u *UPXDownloader) DownloadURL() string {
	if u.apiErr != nil {
		return u.downloadURL()
	}

	asset, err := u.GetLatestReleaseAsset()
	if err != nil {
		return u.downloadURL()
	}

	return asset.GetBrowserDownloadURL()
}

func (u *UPXDownloader) downloadFilename() string {
	return fmt.Sprintf("upx-%s-win64.zip", u.version)
}

func (u *UPXDownloader) DownloadFilename() string {
	if u.apiErr != nil {
		return u.downloadFilename()
	}

	asset, err := u.GetLatestReleaseAsset()
	if err != nil {
		return u.downloadFilename()
	}

	return asset.GetName()
}

func (u *UPXDownloader) Download(path string) error {
	return Download(u.DownloadURL(), path)
}

func (u *UPXDownloader) DownloadToTempDir() (path string, err error) {
	return DownloadToTempDir(u.DownloadURL(), u.DownloadFilename())
}
