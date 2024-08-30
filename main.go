package main

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/server"
	"food-delivery/pkg/logger"
)

// @title Food Delivery API
// @version 1.0
// @description This is a server for the food delivery service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.bearer BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme
// @host localhost:5005
// @BasePath /api/v1
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
