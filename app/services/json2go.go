package services

import (
	"encoding/json/jsontext"

	"github.com/fhluo/json2go/internal/config"
	"github.com/fhluo/json2go/pkg/json2go"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type JSON2Go struct {
	app *application.App
}

func Json2GoService(app *application.App) application.Service {
	return application.NewServiceWithOptions(&JSON2Go{app: app}, application.ServiceOptions{Route: "/json2go"})
}

func (j *JSON2Go) Generate(s string) string {
	if config.OptionsValidJSONBeforeGeneration.Get() && !jsontext.Value(s).IsValid() {
		j.app.Event.Emit("error", "Invalid JSON")
		return ""
	}

	result, err := json2go.Options{
		TypeName:     "T",
		AllCapsWords: config.AllCapsWords.Get(),
	}.Generate(s)

	if err != nil {
		j.app.Event.Emit("error", err.Error())
		return ""
	}

	return result
}
