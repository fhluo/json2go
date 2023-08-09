package version

import (
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/hashicorp/go-multierror"
	"github.com/samber/lo"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"sort"
)

const version = "0.3.2"

func Get() *semver.Version {
	return lo.Must(semver.NewVersion(version))
}

func GetReleasesVersions() ([]*semver.Version, error) {
	resp, err := http.Get("https://github.com/fhluo/json2go/releases")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("failed to close response body", err)
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := regexp.MustCompile(`"/fhluo/json2go/releases/tag/v(\d+\.\d+\.\d+)"`).FindAllStringSubmatch(string(data), -1)

	var errors error
	versions := lo.Map(results, func(result []string, _ int) *semver.Version {
		v, err := semver.NewVersion(result[1])
		if err != nil {
			errors = multierror.Append(errors, err)
		}
		return v
	})
	if errors != nil {
		return nil, errors
	}

	sort.Sort(sort.Reverse(semver.Collection(versions)))
	return versions, nil
}

func GetLatestReleaseVersion() (*semver.Version, error) {
	versions, err := GetReleasesVersions()
	if err != nil {
		return nil, fmt.Errorf("failed to get releases versions: %w", err)
	}

	if len(versions) == 0 {
		return nil, fmt.Errorf("no release found")
	}

	return versions[0], nil
}
