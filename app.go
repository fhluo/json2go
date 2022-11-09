package main

import (
	"bytes"
	"context"
	gen "github.com/dave/jennifer/jen"
	"github.com/fhluo/json2go/pkg/def"
	"github.com/tidwall/gjson"
	"golang.design/x/clipboard"
	"strings"
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

func (a *App) Generate(json string) string {
	if !gjson.Valid(json) {
		return "invalid json"
	}

	file := gen.NewFile("main")
	file.Add(def.Type("T", gjson.Parse(json)))

	buf := new(bytes.Buffer)
	if err := file.Render(buf); err != nil {
		return err.Error()
	}

	return strings.TrimPrefix(buf.String(), "package main\n\n")
}

func (a *App) WriteClipboard(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))
}

func (a *App) ReadClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}
