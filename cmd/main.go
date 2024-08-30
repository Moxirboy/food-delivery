package cmd

import (
	"food-delivery/internal/configs"
	"food-delivery/internal/server"
	"food-delivery/pkg/logger"
)

//	@title			Food Delivery API
//	@version		1.0
//	@description	This is a server for the food delivery service.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:5005
// @BasePath					/api/v1
//
// @schemes					http https
// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
// @description				 security accessToken. Please add it in the format "AccessToken" to authorize your requests.
func main2() {
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
