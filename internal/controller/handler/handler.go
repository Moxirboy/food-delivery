package handler

import (
	"food-delivery/internal/configs"
	v1 "food-delivery/internal/controller/handler/v1"
	"food-delivery/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetUp(
	g *gin.RouterGroup,
	config *configs.Config,
	log logger.Logger,
) {
	SetUpHandlerV1(
		g.Group("/api/v1"),
		config,
		log,
	)
}

func SetUpHandlerV1(
	group *gin.RouterGroup,
	config *configs.Config,
	log logger.Logger,
) {
	v1.NewAuthHandler(
		group, log,
	)

}
