package v1

import (
	"food-delivery/internal/dto"
	"food-delivery/internal/models"
)

func ToProductResponse(
	product models.Product,
) *dto.Product {
	return &dto.Product{
		Name:        product.Name.String,
		Description: product.Description.String,
		Price:       product.Price.Float64,
		Image:       product.Image.String,
	}
}
func toProductListResponse(
	products models.ProductsList,
) *dto.ProductsList {
	list := make([]*dto.Product, len(products.Product))
	for i, product := range products.Product {
		list[i] = ToProductResponse(*product)
	}

	return &dto.ProductsList{
		TotalCount: products.TotalCount,
		TotalPages: products.TotalPages,
		Page:       products.Page,
		Size:       products.Size,
		HasMore:    products.HasMore,
		Product:    list,
	}
}
