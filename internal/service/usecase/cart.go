package usecase

import (
	"context"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type CartUsecase struct {
	log  logger.Logger
	repo repo.Cart
}

func NewCartUsecase(repo repo.Cart, log logger.Logger) ICartUseCase {
	return &CartUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *CartUsecase) Create(ctx context.Context, cart *models.Cart) error {
	if err := uc.repo.Create(ctx, cart); err != nil {
		return err
	}

	return nil
}

func (uc *CartUsecase) GetByID(ctx context.Context, id string) (*models.Cart, error) {
	cart, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (uc *CartUsecase) AddProduct(ctx context.Context, id string, cart *models.CartProduct) error {
	exists, err := uc.repo.CheckStatus(ctx, cart.CartID)
	if err != nil {
		return err
	}
	if !exists {
		Newcart := &models.Cart{
			UserID: id,
			Status: models.CartStatusPending,
		}
		err = uc.repo.Create(ctx, Newcart)
		if err != nil {
			return err
		}
		cart.CartID = Newcart.ID
	}
	return uc.repo.AddProduct(ctx, cart)
}

func (uc *CartUsecase) UpdateQuantity(ctx context.Context, cart *models.CartProduct) error {
	return uc.repo.UpdateQuantity(ctx, cart)
}
