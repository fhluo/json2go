package main

import (
	"embed"
	"github.com/fhluo/json2go/internal/config"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
)

var (
	//go:embed all:web/dist
	assets embed.FS
)

func main() {
	defer config.Write()

	app := NewApp()

	a := application.NewWithOptions(&options.App{
		Title:     "json2go",
		Width:     1200,
		MinWidth:  900,
		Height:    800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Mica,
		},
	})

	app.exit = func() {
		a.Quit()
	}

	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}
}
