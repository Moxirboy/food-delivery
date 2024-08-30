package repo

import (
	"context"
	"food-delivery/internal/models"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	GetOrders(ctx context.Context) ([]*models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) error
	DeleteOrder(ctx context.Context, id string) error
	GetProductOrderDetails(ctx context.Context, id string) (*models.OrderProducts, error)
	GetOrdersByCourierID(ctx context.Context, userID int) ([]*models.Order, error) 
	GetOrdersByStatus(ctx context.Context, status string) ([]*models.Order, error) 
	GetOrderByCartID(ctx context.Context, cartID int) (*models.Order, error)
}
