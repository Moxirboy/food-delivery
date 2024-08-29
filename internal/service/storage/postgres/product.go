package postgres

import (
	"context"
	"database/sql"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type product struct {
	db  *sql.DB
	log logger.Logger
}

func NewProduct(db *sql.DB, log logger.Logger) repo.Product {
	return &product{
		db:  db,
		log: log,
	}
}

func (db *product) Create(ctx context.Context, product *models.Product) error {
	if err := db.db.QueryRowContext(ctx, CreateProduct,
		product.Name, product.Description, product.Price, product.Image,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
		db.log.Error("could not create product:" + err.Error())
		return err
	}

	return nil
}

func (db *product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product := models.Product{}
	if err := db.db.QueryRowContext(ctx, GetProduct, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price,
		&product.Image, &product.CreatedAt, &product.UpdatedAt,
		&product.DeletedAt); err != nil {
		db.log.Error("could not get product by id: ", err.Error())
		return nil, err
	}

	return &product, nil
}

func (db *product) GetList(ctx context.Context) ([]models.Product, error) {
	rows, err := db.db.QueryContext(ctx, GetProductList)
	if err != nil {
		db.log.Error("could not get product list: ", err.Error())
		return nil, err
	}

	response := make([]models.Product)
	for rows.Next() {
	}
}
