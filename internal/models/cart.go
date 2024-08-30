package models

import "database/sql"

type CartStatus string

const (
	CartStatusPending CartStatus = "PENDING"
	CartStatusPaid    CartStatus = "ORDERED"
)


type Cart struct {
	ID     string
	UserID string
	Status CartStatus
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type CartProduct struct {
	CartID    string
	ProductID string
	Quantity  int
}

func NewCartProduct(cartID, productID string, quantity int) *CartProduct {
	return &CartProduct{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
}
func NewCart(userID string) *Cart {
	return &Cart{
		UserID: userID,
		Status: CartStatusPending,
	}
}