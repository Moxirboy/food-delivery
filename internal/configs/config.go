package configs

import (
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var instance Config

func Load() *Config {
	if err := env.Parse(&instance); err != nil {
		panic(err)
	}

	return &instance
}

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`

	Server   Server
	Logger   Logger
	Postgres Postgres
	JWT      JWT
	Setup    Setup
}

type (
	Server struct {
		Environment       string `env:"SERVER_ENVIRONMENT"`
		Port              uint16 `env:"ADMIN_PORT"`
		MaxConnectionIdle uint16 `env:"SERVER_MAX_CONNECTION_IDLE"`
		Timeout           uint16 `env:"SERVER_TIMEOUT"`
		Time              uint16 `env:"SERVER_TIME"`
		MaxConnectionAge  uint16 `env:"SERVER_MAX_CONNECTION_AGE"`
	}

	Setup struct {
		AdminName     string `env:"SETUP_ADMIN_NAME"`
		AdminLastName string `env:"SETUP_ADMIN_LAST_NAME"`
		AdminEmail    string `env:"SETUP_ADMIN_PHONE"`
		AdminPassword string `env:"SETUP_ADMIN_PASSWORD"`
	}

	Logger struct {
		Level    string `env:"LOGGER_LEVEL"`
		Encoding string `env:"LOGGER_ENCODING"`
	}

	Postgres struct {
		Port     uint16 `env:"POSTGRES_PORT"`
		Host     string `env:"POSTGRES_HOST"`
		Password string `env:"POSTGRES_PASSWORD"`
		User     string `env:"POSTGRES_USER"`
		Database string `env:"POSTGRES_DB"`
	}

	JWT struct {
		SecretKeyExpireMinutes   uint16 `env:"JWT_SECRET_KEY_EXPIRE_MINUTES_ADMIN"`
		SecretKey                string `env:"JWT_SECRET_KEY_ADMIN"`
		RefreshKeyExpireHours    uint16 `env:"JWT_REFRESH_KEY_EXPIRE_HOURS_ADMIN"`
		ClientRefreshExpireHours uint16 `env:"JWT_CLIENT_REFRESH_EXPIRE_HOURS"`
		RefreshKey               string `env:"JWT_REFRESH_KEY_ADMIN"`
	}
)
