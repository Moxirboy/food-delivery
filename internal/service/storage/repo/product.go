package repo

import (
	"context"
	"food-delivery/internal/models"
	"food-delivery/pkg/utils"
)

type Product interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error)
	Delete(ctx context.Context, id string) error
}
