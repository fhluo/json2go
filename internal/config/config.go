package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var path = filepath.Join(os.Getenv("LocalAppData"), "json2go")

func init() {
	viper.SetDefault("locale", "")
	viper.SetDefault("font_size", 16)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfigFile()
		}
	}
}

func createConfigFile() {
	if err := os.MkdirAll(path, 0666); err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(filepath.Join(path, "config.toml"))
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

type config struct {
	sync.Mutex
}

func (c *config) Set(key string, value any) {
	c.Lock()
	defer c.Unlock()
	viper.Set(key, value)
}

func (c *config) Get(key string) any {
	c.Lock()
	defer c.Unlock()
	return viper.Get(key)
}

var (
	c   config
	Set = c.Set
	Get = c.Get
)
