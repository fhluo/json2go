package services

import (
	"github.com/fhluo/json2go/internal/examples"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Examples struct{}

func ExamplesService() application.Service {
	return application.NewService(&Examples{}, application.ServiceOptions{Route: "/examples"})
}

func (e *Examples) All() []examples.Example {
	return examples.Examples
}
