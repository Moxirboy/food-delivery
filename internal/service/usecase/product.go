package usecase

import (
	"context"
	"fmt"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
)

type ProductUsecase struct {
	log  logger.Logger
	repo repo.Product
}

func NewProductUsecase(repo repo.Product, log logger.Logger) IProductUseCase {
	return &ProductUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ProductUsecase) Create(ctx context.Context, product *models.Product) error {
	fmt.Println(product.Price)
	if product.Price.Float64 <= 0 {
		// TODO: handle error
		return utils.ErrInvalidField
	}

	if err := uc.repo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}

func (uc *ProductUsecase) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUsecase) GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error) {
	return uc.repo.GetList(
		ctx,
		name,
		query,
	)
}

func (uc *ProductUsecase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
