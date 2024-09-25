package config

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var configFilename string

func Init(filename string) {
	configFilename = filename
	path := filepath.Dir(configFilename)

	// 创建配置文件目录
	err := os.MkdirAll(path, 0666)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// 添加配置文件目录，设置配置文件名
	viper.AddConfigPath(path)

	base := filepath.Base(filename)
	ext := filepath.Ext(base)

	viper.SetConfigName(strings.TrimSuffix(base, ext))
	viper.SetConfigType(strings.TrimPrefix(ext, "."))

	// 读取配置文件
	err = viper.ReadInConfig()
	if err == nil {
		return
	}

	// 读取配置文件失败，判断错误类型
	var notFoundErr viper.ConfigFileNotFoundError
	if !errors.As(err, &notFoundErr) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// 未找到配置文件，创建配置文件
	slog.Info(notFoundErr.Error())
	if err = viper.WriteConfigAs(configFilename); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

// Save 保存配置文件
func Save() {
	// 写入配置
	err := viper.WriteConfig()
	if err == nil {
		return
	}

	// 写入配置失败，判断错误类型
	var notFoundErr viper.ConfigFileNotFoundError
	if !errors.As(err, &notFoundErr) {
		slog.Warn(err.Error())
		return
	}

	// 重新尝试写入配置
	if err = viper.WriteConfigAs(configFilename); err != nil {
		slog.Warn(err.Error())
	}
}
