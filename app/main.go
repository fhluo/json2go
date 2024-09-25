package main

import (
	"embed"
	"github.com/fhluo/json2go/app/services"
	"github.com/fhluo/json2go/internal/config"
	"github.com/lmittmann/tint"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
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

	app := application.New(application.Options{
		Name:        "json2go",
		Description: "",
		Services: []application.Service{
			services.ClipboardService(),
			services.ConfigService(),
			services.DialogsService(),
			services.ExamplesService(),
			services.Json2GoService(),
			services.VersionService(),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Logger: slog.New(
			tint.NewHandler(os.Stderr, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	})

	window := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "json2go",
		Width:     800,
		Height:    600,
		MinWidth:  800,
		MinHeight: 600,

		BackgroundType:   application.BackgroundTypeTranslucent,
		BackgroundColour: application.NewRGBA(27, 38, 54, 1),
		Windows: application.WindowsWindow{
			BackdropType: application.Mica,
		},
	})

	app.OnEvent("exit", func(event *application.CustomEvent) {
		app.Quit()
	})

	app.OnEvent("resize", func(event *application.CustomEvent) {
		config.SetWindowSize(window.Size())
	})

	if err := app.Run(); err != nil {
		slog.Error(err.Error())
	}
}
