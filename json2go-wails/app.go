package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"strings"

	"github.com/fhluo/json2go/json2go"
	"github.com/fhluo/json2go/json2go-wails/i18n"
	"github.com/fhluo/json2go/json2go-wails/internal/config"
	"github.com/fhluo/json2go/json2go-wails/internal/examples"
	"github.com/fhluo/json2go/json2go-wails/internal/version"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

type App struct {
	ctx  context.Context
	exit func()
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	runtime.EventsOn(ctx, "exit", func(optionalData ...interface{}) {
		if a.exit != nil {
			a.exit()
		}
	})

	runtime.EventsOn(ctx, "resize", func(optionalData ...interface{}) {
		config.SetWindowSize(runtime.WindowGetSize(ctx))
	})
}

func (a *App) GetLocale() string {
	slog.Debug("GetLocale", "locale", config.Locale.Get())
	return config.Locale.Get()
}

func (a *App) SetLocale(locale string) {
	slog.Debug("SetLocale", "locale", locale)
	config.Locale.Set(locale)
}

func (a *App) GetFontSize() float64 {
	slog.Debug("GetFontSize", "size", config.FontSize.Get())
	return config.FontSize.Get()
}

func (a *App) SetFontSize(size float64) {
	slog.Debug("SetFontSize", "size", size)
	config.FontSize.Set(size)
}

func (a *App) GetAllCapsWords() []string {
	slog.Debug("GetAllCapsWords", "words", config.AllCapsWords.Get())
	return config.AllCapsWords.Get()
}

func (a *App) SetAllCapsWords(words []string) {
	slog.Debug("SetAllCapsWords", "words", config.AllCapsWords.Get())
	config.AllCapsWords.Set(words)
}

func (a *App) GetOptionsValidJSONBeforeGeneration() bool {
	return config.OptionsValidJSONBeforeGeneration.Get()
}

func (a *App) SetOptionsValidJSONBeforeGeneration(valid bool) {
	config.OptionsValidJSONBeforeGeneration.Set(valid)
}

func (a *App) GetOptionsGenerateInRealTime() bool {
	return config.OptionsGenerateInRealTime.Get()
}

func (a *App) SetOptionsGenerateInRealTime(b bool) {
	config.OptionsGenerateInRealTime.Set(b)
}

func (a *App) GetVersion() string {
	return version.Get().String()
}

func (a *App) GetLatestVersion() string {
	v, err := version.GetLatestReleaseVersion()
	if err != nil {
		return ""
	}

	return v.String()
}

func (a *App) GetExamples() []examples.Example {
	return examples.Examples
}

func (a *App) Generate(s string) string {
	if a.GetOptionsValidJSONBeforeGeneration() && !json.Valid([]byte(s)) {
		runtime.EventsEmit(a.ctx, "error", "invalid json")
		return ""
	}

	result, err := json2go.Options{
		TypeName:     "T",
		AllCapsWords: a.GetAllCapsWords(),
	}.Generate(s)

	if err != nil {
		runtime.EventsEmit(a.ctx, "error", err.Error())
		return ""
	}

	return result
}

func (a *App) OpenJSONFile() string {
	filename, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: i18n.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Open JSON File",
				Other: "Open JSON File",
			},
		}),
		CanCreateDirectories: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: i18n.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: &i18n.Message{
						ID:    "JSON Files(*.json)",
						Other: "JSON Files(*.json)",
					},
				}),
				Pattern: "*.json",
			},
		},
	})
	if err != nil {
		return ""
	}

	// cancelled by user
	if filename == "" {
		return ""
	}

	// read file and return string content
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(data)
}

func (a *App) SaveGoSourceFile(s string) {
	filename, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: i18n.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Save Go Source File",
				Other: "Save Go Source File",
			},
		}),
		CanCreateDirectories: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: i18n.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: &i18n.Message{
						ID:    "Go Source Files(*.go)",
						Other: "Go Source Files(*.go)",
					},
				}),
				Pattern: "*.go",
			},
		},
	})
	if err != nil {
		return
	}

	// cancelled by user
	if filename == "" {
		return
	}

	// if filename doesn't end with .go, append it
	if !strings.HasSuffix(filename, ".go") {
		filename += ".go"
	}

	// write file
	_ = os.WriteFile(filename, []byte(s), 0664)
}

func (a *App) WriteClipboard(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))
}

func (a *App) ReadClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}
