package usecase

import (
	"context"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type orderUseCase struct {
	orderRepo repo.IOrderRepository
	cartRepo  repo.Cart
	log       logger.Logger
}

func NewOrderUseCase(
	orderRepo repo.IOrderRepository,
	log logger.Logger,
) IOrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
		log:       log,
	}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, order models.Order) error {
	err := uc.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		uc.log.Errorf("Error while creating order: %v", err)
		return err
	}

	err = uc.cartRepo.UpdateStatus(ctx, order.ID, models.CartStatusPaid)
	if err != nil {
		uc.log.Errorf("Error while creating order: %v", err)
		return err
	}
	return nil
}

func (uc *orderUseCase) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return uc.orderRepo.GetOrder(ctx, id)
}

func (uc *orderUseCase) GetOrders(ctx context.Context) ([]*models.Order, error) {
	return uc.orderRepo.GetOrders(ctx)
}

func (uc *orderUseCase) UpdateOrder(ctx context.Context, order models.Order) error {
	return uc.orderRepo.UpdateOrder(ctx, order)
}

func (uc *orderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return uc.orderRepo.DeleteOrder(ctx, id)
}
