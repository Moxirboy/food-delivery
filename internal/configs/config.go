package configs

import (
	"os"
	"strconv"
)

//import (
//	env "github.com/caarlos0/env/v6"
//	_ "github.com/joho/godotenv/autoload"
//)
//
//var instance Config
//
//func Load() *Config {
//	if err := env.Parse(&instance); err != nil {
//		panic(err)
//	}
//
//	return &instance
//}

var instance Config

func Load() *Config {
	instance = Config{
		AppName:    getEnv("APP_NAME", ""),
		AppVersion: getEnv("APP_VERSION", ""),

		Server: Server{
			Environment:       getEnv("SERVER_ENVIRONMENT", ""),
			Port:              parseUint16(getEnv("ADMIN_PORT", "0")),
			MaxConnectionIdle: parseUint16(getEnv("SERVER_MAX_CONNECTION_IDLE", "0")),
			Timeout:           parseUint16(getEnv("SERVER_TIMEOUT", "0")),
			Time:              parseUint16(getEnv("SERVER_TIME", "0")),
			MaxConnectionAge:  parseUint16(getEnv("SERVER_MAX_CONNECTION_AGE", "0")),
		},

		Logger: Logger{
			Level:    getEnv("LOGGER_LEVEL", ""),
			Encoding: getEnv("LOGGER_ENCODING", ""),
		},

		Postgres: Postgres{
			Port:     parseUint16(getEnv("POSTGRES_PORT", "0")),
			Host:     getEnv("POSTGRES_HOST", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			User:     getEnv("POSTGRES_USER", ""),
			Database: getEnv("POSTGRES_DATABASE", ""),
		},

		JWT: JWT{
			SecretKeyExpireMinutes:   parseUint16(getEnv("JWT_SECRET_KEY_EXPIRE_MINUTES_ADMIN", "0")),
			SecretKey:                getEnv("JWT_SECRET_KEY_ADMIN", ""),
			RefreshKeyExpireHours:    parseUint16(getEnv("JWT_REFRESH_KEY_EXPIRE_HOURS_ADMIN", "0")),
			ClientRefreshExpireHours: parseUint16(getEnv("JWT_CLIENT_REFRESH_EXPIRE_HOURS", "0")),
			RefreshKey:               getEnv("JWT_REFRESH_KEY_ADMIN", ""),
		},

		Casbin: Casbin{
			ConfigPath: getEnv("CASBIN_CONFIG_PATH_ADMIN", ""),
			Name:       getEnv("CASBIN_NAME_ADMIN", ""),
		},

		FireStore: FireStore{
			ProjectID:       getEnv("PROJECT_ID", ""),
			CredentialsFile: getEnv("CREDENTIALS_FILE", ""),
			BucketName:      getEnv("STORAGE_BUCKET_NAME", ""),
		},
	}

	return &instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func parseUint16(value string) uint16 {
	if v, err := strconv.ParseUint(value, 10, 16); err == nil {
		return uint16(v)
	}
	return 0
}

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`

	Server    Server
	Logger    Logger
	Postgres  Postgres
	JWT       JWT
	Setup     Setup
	Redis     Redis
	Casbin    Casbin
	FireStore FireStore
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

	Redis struct {
		Host     string `env:"REDIS_HOST"`
		Port     uint16 `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
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
		Database string `env:"POSTGRES_DATABASE"`
	}

	JWT struct {
		SecretKeyExpireMinutes   uint16 `env:"JWT_SECRET_KEY_EXPIRE_MINUTES_ADMIN"`
		SecretKey                string `env:"JWT_SECRET_KEY_ADMIN"`
		RefreshKeyExpireHours    uint16 `env:"JWT_REFRESH_KEY_EXPIRE_HOURS_ADMIN"`
		ClientRefreshExpireHours uint16 `env:"JWT_CLIENT_REFRESH_EXPIRE_HOURS"`
		RefreshKey               string `env:"JWT_REFRESH_KEY_ADMIN"`
	}
	Casbin struct {
		ConfigPath string `env:"CASBIN_CONFIG_PATH_ADMIN"`
		Name       string `env:"CASBIN_NAME_ADMIN"`
	}

	FireStore struct {
		ProjectID       string `env:"PROJECT_ID"`
		CredentialsFile string `env:"CREDENTIALS_FILE"`
		BucketName      string `env:"STORAGE_BUCKET_NAME"`
	}
)
