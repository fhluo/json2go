package services

import (
	"encoding/json"
	"github.com/fhluo/json2go/internal/config"
	"github.com/fhluo/json2go/pkg/json2go"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type JSON2Go struct{}

func Json2GoService() application.Service {
	return application.NewService(&JSON2Go{}, application.ServiceOptions{Route: "/json2go"})
}

func (j *JSON2Go) Generate(s string) string {
	if config.OptionsValidJSONBeforeGeneration.Get() && !json.Valid([]byte(s)) {
		//runtime.EventsEmit(j.ctx, "error", "invalid json")
		return ""
	}

	result, err := json2go.Options{
		TypeName:     "T",
		AllCapsWords: config.AllCapsWords.Get(),
	}.Generate(s)

	if err != nil {
		//runtime.EventsEmit(j.ctx, "error", err.Error())
		return ""
	}

	return result
}
