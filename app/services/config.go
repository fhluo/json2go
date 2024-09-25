package services

import (
	"github.com/fhluo/json2go/internal/config"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log/slog"
)

type Config struct{}

func ConfigService() application.Service {
	return application.NewService(&Config{}, application.ServiceOptions{Route: "/config"})
}

func (c *Config) GetLocale() string {
	slog.Debug("GetLocale", "locale", config.Locale.Get())
	return config.Locale.Get()
}

func (c *Config) SetLocale(locale string) {
	slog.Debug("SetLocale", "locale", locale)
	config.Locale.Set(locale)
}

func (c *Config) GetFontSize() float64 {
	slog.Debug("GetFontSize", "size", config.FontSize.Get())
	return config.FontSize.Get()
}

func (c *Config) SetFontSize(size float64) {
	slog.Debug("SetFontSize", "size", size)
	config.FontSize.Set(size)
}

func (c *Config) GetAllCapsWords() []string {
	slog.Debug("GetAllCapsWords", "words", config.AllCapsWords.Get())
	return config.AllCapsWords.Get()
}

func (c *Config) SetAllCapsWords(words []string) {
	slog.Debug("SetAllCapsWords", "words", config.AllCapsWords.Get())
	config.AllCapsWords.Set(words)
}

func (c *Config) GetOptionsValidJSONBeforeGeneration() bool {
	return config.OptionsValidJSONBeforeGeneration.Get()
}

func (c *Config) SetOptionsValidJSONBeforeGeneration(valid bool) {
	config.OptionsValidJSONBeforeGeneration.Set(valid)
}

func (c *Config) GetOptionsGenerateInRealTime() bool {
	return config.OptionsGenerateInRealTime.Get()
}

func (c *Config) SetOptionsGenerateInRealTime(b bool) {
	config.OptionsGenerateInRealTime.Set(b)
}
