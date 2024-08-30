package dto


type OrderResponse struct {
    ID     string `json:"id"`
    CartID string `json:"cart_id"`
    Status string `json:"status"`
}

type OrderStatus struct{
	Status string `json:"status"`
}