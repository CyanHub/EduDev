package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Logging struct {
		Level string
		Path  string
	}
	Server struct {
		Port      int
		JWTSecret string
	}
}

var (
	once     sync.Once
	instance *Config
)

func Load() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		instance = &Config{}
		if err := viper.Unmarshal(instance); err != nil {
			log.Fatalf("Unable to decode into struct: %v", err)
		}
	})
	return instance
}
