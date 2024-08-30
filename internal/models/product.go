package models

import (
	"database/sql"
	"food-delivery/internal/dto"
)

type Product struct {
	ID          string
	Name        sql.NullString
	Description sql.NullString
	Price       sql.NullFloat64
	Image       sql.NullString

	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type ProductsList struct {
	TotalCount int
	TotalPages int
	Page       int
	Size       int
	HasMore    bool
	Product    []*Product
}

func NewProduct(model dto.Product) *Product {
	return &Product{
		Name:        sql.NullString{Valid: model.Name != "", String: model.Name},
		Description: sql.NullString{Valid: model.Description != "", String: model.Description},
		Price:       sql.NullFloat64{Valid: model.Price != 0, Float64: model.Price},
		Image:       sql.NullString{Valid: model.Image != "", String: model.Image},
	}
}
