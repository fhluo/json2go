package main

import (
	"embed"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/fhluo/json2go/json2go-wails/internal/config"
	"github.com/fhluo/json2go/json2go-wails/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

var (
	//go:embed all:web/dist
	assets  embed.FS
	logFile *os.File
)

func init() {
	log.SetFlags(0)

	var err error

	logFile = os.Stderr
	if os.Getenv("json2go_dev") != "true" {
		logFile, err = os.OpenFile(filepath.Join(config.Path, "log.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}

	level := slog.LevelInfo
	if os.Getenv("json2go_debug") == "true" {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: level,
	})))
}

func main() {
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Println(err)
		}
		config.Save()
	}()

	app := NewApp()

	a := application.NewWithOptions(&options.App{
		Title:     "json2go",
		Width:     800,
		MinWidth:  800,
		Height:    600,
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
		Logger: logger.New(logFile),
	})

	app.exit = func() {
		a.Quit()
	}

	if err := a.Run(); err != nil {
		slog.Error(err.Error())
	}
}
