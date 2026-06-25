package services

import (
	"github.com/fhluo/json2go/internal/version"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Version struct{}

func VersionService() application.Service {
	return application.NewServiceWithOptions(&Version{}, application.ServiceOptions{Route: "/version"})
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

type UpdateInfo struct {
	CurrentVersion string `json:"currentVersion"`
	LatestVersion  string `json:"latestVersion"`
	HasUpdate      bool   `json:"hasUpdate"`
}

func (v *Version) CheckForUpdate() UpdateInfo {
	current := version.Version()
	latest, err := version.Latest()
	if err != nil {
		return UpdateInfo{}
	}

	return UpdateInfo{
		CurrentVersion: current.String(),
		LatestVersion:  latest.String(),
		HasUpdate:      current.LessThan(latest),
	}
}
