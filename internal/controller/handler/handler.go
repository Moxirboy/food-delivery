package handler

import (
	_ "food-delivery/docs"
	"food-delivery/internal/configs"
	v1 "food-delivery/internal/controller/handler/v1"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp(
	g *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
	redis repo.IRedisRepository,
) {
	SetUpHandlerV1(
		g.Group("/api/v1"),
		config, uc, log,
	)
	url := ginSwagger.URL("swagger/doc.json")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func SetUpHandlerV1(
	group *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
) {
	v1.NewAuthHandler(
		group,
		log,
		uc.IAuthUseCase(),
		uc.IAccountUseCase(),
	)
	v1.NewProductHandler(
		group,
		log,
		config,
		uc.ProductUsecase(),
	)
	v1.NewCartHandler(
		group,
		log,
		uc.CartUsecase(),
	)
	v1.NeworderHandler(
		group,
		log,
		uc.IOrderUseCase(),
		uc.CartUsecase(),
	)
}
