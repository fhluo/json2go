package i18n

import (
	"embed"
	"github.com/fhluo/json2go/internal/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"github.com/samber/lo"
	"golang.org/x/sys/windows"
	"golang.org/x/text/language"
	"io/fs"
	"log"
	"strings"
)

type (
	LocalizeConfig = i18n.LocalizeConfig
	Message        = i18n.Message
)

var (
	//go:embed locales
	locales embed.FS

	localizer *i18n.Localizer

	MustLocalize func(lc *LocalizeConfig) string
)

//go:generate goi18n extract -outdir locales ../
//go:generate goi18n merge -outdir locales locales/active.en.toml locales/translate.zh-Hans.toml
//go:generate goi18n merge -outdir locales locales/active.en.toml locales/active.zh-Hans.toml
//go:generate goi18n merge -outdir locales locales/active.en.toml locales/active.zh-Hans.toml locales/translate.zh-Hans.toml

func init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// load locales
	err := fs.WalkDir(locales, "locales", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasPrefix(d.Name(), "active.") {
			lo.Must(bundle.LoadMessageFileFS(locales, path))
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	var languages []string

	lang := config.GetLocale()
	if lang != "" {
		languages = append(languages, lang)
	} else {
		languages, err = windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
		if err != nil {
			languages, err = windows.GetSystemPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	localizer = i18n.NewLocalizer(bundle, languages...)
	MustLocalize = localizer.MustLocalize
}
