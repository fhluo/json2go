package version

import (
	_ "embed"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io"
	"iter"
	"log/slog"
	"net/http"
	"regexp"
	"slices"
	"unsafe"

	"github.com/Masterminds/semver"
)

//go:embed version.toml
var versionTOML []byte

type versionInfo struct {
	Version    string `toml:"version"`
	ReleaseURL string `toml:"release_url"`
}

var info versionInfo

func init() {
	if err := toml.Unmarshal(versionTOML, &info); err != nil {
		slog.Error("failed to unmarshal version toml", "err", err)
	}
}

func Version() *semver.Version {
	v, err := semver.NewVersion(info.Version)

	if err != nil {
		slog.Error("failed to parse version", "err", err)
		return nil
	}

	return v
}

func fetchReleasePage() (string, error) {
	resp, err := http.Get(info.ReleaseURL)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", "err", err)
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(data), len(data)), nil
}

func All() iter.Seq2[*semver.Version, error] {
	return func(yield func(*semver.Version, error) bool) {
		page, err := fetchReleasePage()
		if err != nil {
			yield(nil, err)
			return
		}

		results := regexp.MustCompile(`"/fhluo/json2go/releases/tag/v(\d+\.\d+\.\d+)"`).FindAllStringSubmatch(page, -1)

		for result := range slices.Values(results) {
			if !yield(semver.NewVersion(result[1])) {
				return
			}
		}
	}
}

func collect[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var items []T

	for item, err := range seq {
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func Latest() (*semver.Version, error) {
	versions, err := collect(All())
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return nil, fmt.Errorf("no release found")
	}

	slices.SortFunc(versions, (*semver.Version).Compare)
	return versions[len(versions)-1], nil
}
