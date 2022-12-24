package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	keyLocale       = "locale"
	keyFontSize     = "font_size"
	keyWindowWidth  = "window.width"
	keyWindowHeight = "window.height"
	keyAllCapsWords = "all_caps_words"
)

var (
	Path = filepath.Join(os.Getenv("LocalAppData"), "json2go")
	m    = new(sync.Mutex)
)

func init() {
	viper.SetDefault(keyLocale, "")
	viper.SetDefault(keyFontSize, 16)
	viper.SetDefault(keyWindowWidth, 1200)
	viper.SetDefault(keyWindowHeight, 800)
	viper.SetDefault(keyAllCapsWords, []string{"ID", "URL", "URI", "JSON", "HTML", "CSS", "API", "HTTP", "SQL"})

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath(Path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfigFile()
		}
	}
}

func createConfigFile() {
	if err := os.MkdirAll(Path, 0666); err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(filepath.Join(Path, "config.toml"))
	if err != nil {
		log.Fatalln(err)
	}

	if err = f.Close(); err != nil {
		log.Println(err)
	}
}

func Write() {
	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalln(err)
		} else {
			log.Println(err)
		}
	}
}

func set(key string, value any) {
	m.Lock()
	defer m.Unlock()
	viper.Set(key, value)
}

func GetLocale() string {
	m.Lock()
	defer m.Unlock()
	return viper.GetString(keyLocale)
}

func SetLocale(locale string) {
	set(keyLocale, locale)
}

func GetFontSize() float64 {
	m.Lock()
	defer m.Unlock()
	return viper.GetFloat64(keyFontSize)
}

func SetFontSize(size float64) {
	set(keyFontSize, size)
}

func GetWindowSize() (int, int) {
	m.Lock()
	defer m.Unlock()
	return viper.GetInt(keyWindowWidth), viper.GetInt(keyWindowHeight)
}

func SetWindowSize(width, height int) {
	m.Lock()
	defer m.Unlock()
	viper.Set(keyWindowWidth, width)
	viper.Set(keyWindowHeight, height)
}

func GetWindowWidth() int {
	m.Lock()
	defer m.Unlock()
	return viper.GetInt(keyWindowWidth)
}

func SetWindowWidth(width int) {
	set(keyWindowWidth, width)
}

func GetWindowHeight() int {
	m.Lock()
	defer m.Unlock()
	return viper.GetInt(keyWindowHeight)
}

func SetWindowHeight(height int) {
	set(keyWindowHeight, height)
}

func GetAllCapsWords() []string {
	m.Lock()
	defer m.Unlock()
	return viper.GetStringSlice(keyAllCapsWords)
}

func SetAllCapsWords(words []string) {
	set(keyAllCapsWords, words)
}
