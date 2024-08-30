package models

import "errors"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusDelivering OrderStatus = "delivering"
)

// Order represents the order model
type Order struct {
	ID     string
	CartID string
	Status OrderStatus
}

type OrderProducts struct {
	ID      string
	OrderID string
	Status  OrderStatus
	Product []Product
}

func NewOrder(id string, cartID string, status OrderStatus) *Order {
	return &Order{
		ID:     id,
		CartID: cartID,
		Status: status,
	}
}
func StatusMaker(status string) (OrderStatus, error) {
	switch status {
	case string(OrderStatusPending):
		return OrderStatusPending, nil
	case string(OrderStatusCompleted):
		return OrderStatusCompleted, nil
	case string(OrderStatusCancelled):
		return OrderStatusCancelled, nil
	case string(OrderStatusDelivering):
		return OrderStatusDelivering, nil
	default:
		return "", errors.New("invalid order status")
	}
}
