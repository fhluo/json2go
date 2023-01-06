package main

import (
	"embed"
	"github.com/fhluo/json2go/internal/config"
	"github.com/fhluo/json2go/internal/logger"
	"github.com/mattn/go-isatty"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"os"
	"path/filepath"
)

var (
	//go:embed all:web/dist
	assets embed.FS
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer config.Write()

	app := NewApp()

	out := os.Stderr
	if !isatty.IsTerminal(os.Stderr.Fd()) {
		file, err := os.OpenFile(filepath.Join(config.Path, "log.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Println(err)
			}
		}()

		out = file
	}

	a := application.NewWithOptions(&options.App{
		Title:     "json2go",
		Width:     config.GetWindowWidth(),
		MinWidth:  800,
		Height:    config.GetWindowHeight(),
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
		Logger: logger.New(out),
	})

	app.exit = func() {
		a.Quit()
	}

	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}
}
