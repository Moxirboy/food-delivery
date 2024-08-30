package dto

type Product struct {
	Name        string  `json:"name" `
	Description string  `json:"description" `
	Price       float64 `json:"price" `
	Image       string  `json:"image" `
}

type GetProductByID struct {
	ID string `query:"id" binding:"required"`
}
type ProductsList struct {
	TotalCount int        `json:"totalCount" :"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	HasMore    bool       `json:"has_more"`
	Product    []*Product `json:"product"`
}
