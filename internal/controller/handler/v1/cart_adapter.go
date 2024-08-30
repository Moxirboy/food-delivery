package v1

import (
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
)



func ToCartResponse(
	cart models.Cart,
) *dto.CartResponse {
	return &dto.CartResponse{
		ID:     cart.ID,
		UserID: cart.UserID,
		Status: string(cart.Status),
	}
}