package repo

import (
	"context"
	"food-delivery/internal/models"
)

type Product interface {
	Create(ctx context.Context, product *models.Product) error
}
