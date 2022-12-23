package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	Path = filepath.Join(os.Getenv("LocalAppData"), "json2go")
	m    = new(sync.Mutex)
)

const (
	keyLocale   = "locale"
	keyFontSize = "font_size"
)

func init() {
	viper.SetDefault(keyLocale, "")
	viper.SetDefault(keyFontSize, 16)

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
