package services

import (
	"github.com/fhluo/json2go/internal/version"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Version struct{}

func VersionService() application.Service {
	return application.NewService(&Version{}, application.ServiceOptions{Route: "/version"})
}

func (v *Version) GetVersion() string {
	return version.Version().String()
}

func (v *Version) GetLatestVersion() string {
	r, err := version.Latest()
	if err != nil {
		return ""
	}

	return r.String()
}
