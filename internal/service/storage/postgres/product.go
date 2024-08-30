package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
	"github.com/pkg/errors"
	"time"
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

func (r *product) Create(ctx context.Context, product *models.Product) error {
	fmt.Println(product.Image)
	if err := r.db.QueryRowContext(ctx, CreateProduct,
		product.Name, product.Description, product.Price, product.Image,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
		r.log.Error("could not create product:" + err.Error())
		return err
	}

	return nil
}

func (r *product) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product := models.Product{}

	if err := r.db.QueryRowContext(ctx, GetProduct, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price,
		&product.Image, &product.CreatedAt, &product.UpdatedAt,
		&product.DeletedAt); err != nil {
		r.log.Error("could not get product by id: ", err.Error())
		return nil, err
	}

	return &product, nil
}

func (r *product) GetList(ctx context.Context, name string, query utils.PaginationQuery) (*models.ProductsList, error) {
	var (
		totalCount     int
		GetProductList = `
	select
		id, name, description, price, image, created_at, updated_at, deleted_at
	from products
	where deleted_at is null and 1=1
`
		GetTotalCount = `
SELECT COUNT(id) FROM products WHERE 1=1
`
	)
	if name != "" {
		GetTotalCount = fmt.Sprintf("%s%s", GetTotalCount, " and name LIKE '%"+name+"%'")
		GetProductList = fmt.Sprintf("%s%s", GetProductList, " and name LIKE '%"+name+"%'")
	}
	GetProductList += " ORDER BY created_at OFFSET $1 LIMIT $2"
	if err := r.db.QueryRowContext(ctx, GetTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	if totalCount == 0 {
		return &models.ProductsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Product:    make([]*models.Product, 0),
		}, nil
	}
	rows, err := r.db.QueryContext(ctx, GetProductList, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "product.GetAll.QueryContext")
	}
	defer rows.Close()
	productList := make([]*models.Product, 0, query.GetSize())
	for rows.Next() {
		product := &models.Product{}
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		); err != nil {
			return nil, errors.Wrap(err, "product.GetAll.Scan")
		}
		productList = append(productList, product)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "product.GetAll.Row.Err")
	}
	return &models.ProductsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Product:    productList,
	}, nil
}

func (r *product) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, DeleteProduct, time.Now().Format(time.RFC3339), id)
	if err != nil {
		r.log.Error("could not delete product: ", err.Error())
		return err
	}
	return nil
}
