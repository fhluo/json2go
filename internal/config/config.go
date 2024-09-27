package config

import (
	"github.com/fhluo/json2go/pkg/config"
	"os"
	"path/filepath"
)

var (
	Path = filepath.Join(os.Getenv("LocalAppData"), "json2go")

	Locale       = config.NewItem("locale", "")
	FontSize     = config.NewItem("font_size", 16.0)
	WindowWidth  = config.NewItem("window.width", 1200)
	WindowHeight = config.NewItem("window.height", 800)
	AllCapsWords = config.NewItem(
		"all_caps_words",
		[]string{"ID", "URL", "URI", "JSON", "HTML", "CSS", "API", "HTTP", "SQL"},
	)
	OptionsValidJSONBeforeGeneration = config.NewItem("options.valid_json_before_generation", true)
	OptionsGenerateInRealTime        = config.NewItem("options.generate_in_real_time", true)
)

func init() {
	config.Init(filepath.Join(Path, "config.toml"))
}

func GetWindowSize() (int, int) {
	return WindowWidth.Get(), WindowHeight.Get()
}

func SetWindowSize(width, height int) {
	WindowWidth.Set(width)
	WindowHeight.Set(height)
}

func Save() {
	config.Save()
}
