package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fhluo/json2go/pkg/def"
	"golang.design/x/clipboard"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAcronyms() []string {
	return def.GetAcronyms()
}

func (a *App) SetAcronyms(acronyms []string) {
	def.SetAcronyms(acronyms...)
}

func (a *App) Generate(s string) string {
	if !json.Valid([]byte(s)) {
		return "invalid json"
	}

	statement, err := def.From(s).Declare("t")
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
