package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server
		Database *Database
		Jwt      *Jwt
	}

	Server struct {
		Port int
	}

	Database struct {
		Host     string
		Port     uint16
		User     string
		Password string
		DBName   string
		TimeZone string
	}

	Jwt struct {
		Access      string
		Refresh     string
		Access_exp  uint8
		Refresh_exp uint16
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
