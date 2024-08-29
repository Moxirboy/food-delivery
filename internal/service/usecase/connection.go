package usecase

import (
	"database/sql"
	"food-delivery/internal/configs"
	"food-delivery/internal/service/storage/postgres"
	"food-delivery/pkg/logger"
)

type IUseCase interface {
	AuthUsecase() authUsecase
	ProductUsecase() productUsecase
}

type UseCase struct {
	connections map[string]interface{}
}

const (
	_authUseCase    = "auth_use_case"
	_productUseCase = "product_use_case"
)

func New(
	cfg *configs.Config,
	pg *sql.DB,
	logger logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_authUseCase] = NewAuthUsecase(pg, logger)
	connections[_productUseCase] = NewProductUsecase(
		postgres.NewProduct(
			pg,
			logger,
		),
		logger,
	)

	return &UseCase{
		connections: connections,
	}
}

func (c *UseCase) AuthUsecase() authUsecase {
	return c.connections[_authUseCase].(authUsecase)
}

func (c *UseCase) ProductUsecase() productUsecase {
	return c.connections[_productUseCase].(productUsecase)
}
