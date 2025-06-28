package config

import "sync"

type (
	Config struct {
		Server   *Server
		Database *Database
	}

	Server struct {
		port int
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBname   string
		Timezone string
	}
)

var (
	once           sync.Once
	configInstance *Config
)
