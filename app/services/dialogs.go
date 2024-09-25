package services

import (
	"github.com/fhluo/json2go/app/i18n"
	"github.com/wailsapp/wails/v3/pkg/application"
	"os"
	"strings"
)

type Dialogs struct{}

func DialogsService() application.Service {
	return application.NewService(&Dialogs{}, application.ServiceOptions{Route: "/dialogs"})
}

func (dialogs *Dialogs) OpenJSONFile() string {
	dialog := application.OpenFileDialog().SetTitle(i18n.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Open JSON File",
			Other: "Open JSON File",
		},
	})).AddFilter(i18n.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "JSON Files(*.json)",
			Other: "JSON Files(*.json)",
		},
	}), "*.json")

	filename, err := dialog.PromptForSingleSelection()
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

func (dialogs *Dialogs) SaveGoSourceFile(s string) {
	dialog := application.SaveFileDialogWithOptions(&application.SaveFileDialogOptions{
		Title: i18n.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Save Go Source File",
				Other: "Save Go Source File",
			},
		}),
	}).AddFilter(i18n.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Go Source Files(*.go)",
			Other: "Go Source Files(*.go)",
		},
	}), "*.go").CanCreateDirectories(true)

	filename, err := dialog.PromptForSingleSelection()
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
