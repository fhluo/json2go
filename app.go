package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fhluo/json2go/internal/config"
	"github.com/fhluo/json2go/pkg/def"
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

	statement, err := def.From(s, allCaps...).Declare("t")
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)
	if err := statement.Render(buf); err != nil {
		return err.Error()
	}

	return buf.String()
}

func (a *App) WriteClipboard(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))
}

func (a *App) ReadClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}
