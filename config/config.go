package config

import (
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

var (
	once sync.Once
	conf *Config
)

func Get() *Config {
	once.Do(func() {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}

		conf = &Config{Env: env}
		configor.New(&configor.Config{Environment: env}).Load(conf, "config/config.yml")
	})
	return conf
}
