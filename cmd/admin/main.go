package main

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/server"
	"food-delivery/pkg/logger"
)

func main() {
	var (
		config = configs.Load()
	)

	logger := logger.NewLogger(config.Logger.Level, config.Logger.Encoding)
	logger.InitLogger()

	logger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)

	s := server.NewServer(config, logger)

	logger.Fatal(s.Run())
}
