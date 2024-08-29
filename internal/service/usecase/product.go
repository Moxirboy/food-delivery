package usecase

import (
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type ProductUsecase struct {
	log  logger.Logger
	repo repo.Product
}

func NewProductUsecase(repo repo.Product, log logger.Logger) ProductUsecase {
	return ProductUsecase{
		log:  log,
		repo: repo,
	}
}
