package dto


type CartRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}


type CartResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Status string `json:"status"`
}