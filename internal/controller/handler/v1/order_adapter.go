package v1

import (
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
)



func ToOrderResponse(order *models.Order) *dto.OrderResponse {
    return &dto.OrderResponse{
        ID:     order.ID,
        CartID: order.CartID,
        Status: string(order.Status), // Assuming OrderStatus is a custom type, convert it to a string
    }
}


func ToOrderResponseArray(orders []*models.Order) []*dto.OrderResponse {
    responses := make([]*dto.OrderResponse, len(orders))
    for i, order := range orders {
        responses[i] = ToOrderResponse(order)
    }
    return responses
}