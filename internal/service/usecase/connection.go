package usecase

import (
	"database/sql"
	"food-delivery/internal/configs"
	"food-delivery/pkg/logger"
)

type IUseCase interface {
}

type UseCase struct {
	connections map[string]interface{}
}

const ()

func New(
	cfg *configs.Config,
	pg *sql.DB,
	logger logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})

	return &UseCase{
		connections: connections,
	}
}
