package repo

import (
	"context"
	"food-delivery/internal/models"
)

type Cart interface {
	Create(ctx context.Context, cart *models.Cart) error
	GetByID(ctx context.Context, id string) (*models.Cart, error)
	AddProduct(ctx context.Context, cartProduct *models.CartProduct) error
	CheckStatus(ctx context.Context, id string) (bool, error) 
	UpdateStatus(ctx context.Context, id string, status models.CartStatus) error
	UpdateQuantity(ctx context.Context, cart *models.CartProduct) error
}
