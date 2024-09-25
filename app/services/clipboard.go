package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.design/x/clipboard"
	"unsafe"
)

func init() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

type Clipboard struct{}

func ClipboardService() application.Service {
	return application.NewService(&Clipboard{}, application.ServiceOptions{Route: "/clipboard"})
}

func (c *Clipboard) Write(text string) {
	clipboard.Write(clipboard.FmtText, unsafe.Slice(unsafe.StringData(text), len(text)))
}

func (c *Clipboard) Read() string {
	data := clipboard.Read(clipboard.FmtText)
	return unsafe.String(unsafe.SliceData(data), len(data))
}
