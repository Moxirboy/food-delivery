package server

import (
	"fmt"

	"food-delivery/internal/configs"
	"food-delivery/internal/controller/handler"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/postgres"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *configs.Config
	logger logger.Logger
}

func NewServer(
	cfg *configs.Config,
	logger logger.Logger,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) Run() error {

	pDB, err := postgres.DB(&s.cfg.Postgres)
	if err != nil {
		s.logger.Fatal(err)
	}

	g := gin.New()

	uc := usecase.New(s.cfg, pDB, s.logger)
	uc.AuthUsecase().Create()

	handler.SetUp(&g.RouterGroup, s.cfg, s.logger)
	_ = uc

	return g.Run(fmt.Sprintf(":%d", s.cfg.Server.Port))

}
