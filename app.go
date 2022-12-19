package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fhluo/json2go/i18n"
	"github.com/fhluo/json2go/internal/config"
	"github.com/fhluo/json2go/pkg/def"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
	"os"
	"strings"
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
}

func (a *App) SetConfig(key string, value any) {
	config.Set(key, value)
}

func (a *App) GetConfig(key string) any {
	return config.Get(key)
}

func (a *App) Generate(s string, allCaps []string) string {
	if !json.Valid([]byte(s)) {
		return "invalid json"
	}

	statement, err := def.From(s, allCaps...).Declare("T")
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)
	if err = statement.Render(buf); err != nil {
		return err.Error()
	}

	return buf.String()
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
