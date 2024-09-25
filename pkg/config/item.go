package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log/slog"
	"sync"
)

var mutex sync.Mutex

type Item[T any] struct {
	Key          string
	DefaultValue T
}

func NewItem[T any](key string, defaultValue T) Item[T] {
	viper.SetDefault(key, defaultValue)
	return Item[T]{
		Key:          key,
		DefaultValue: defaultValue,
	}
}

func (item Item[T]) Get() T {
	mutex.Lock()
	defer mutex.Unlock()

	value := viper.Get(item.Key)
	switch value.(type) {
	case []interface{}:
		switch any(item.DefaultValue).(type) {
		case []string:
			value = cast.ToStringSlice(value)
		case []int:
			value = cast.ToIntSlice(value)
		}
	}

	r, ok := value.(T)
	if !ok {
		slog.Warn("conversion failure", "type", fmt.Sprintf("%T", value))
		return item.DefaultValue
	}

	return r
}

func (item Item[T]) Set(value T) {
	mutex.Lock()
	defer mutex.Unlock()

	viper.Set(item.Key, value)
}
