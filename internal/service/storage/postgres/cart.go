package postgres

import (
	"context"
	"database/sql"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type cart struct {
	db  *sql.DB
	log logger.Logger
}

func NewCart(db *sql.DB, log logger.Logger) repo.Cart {
	return &cart{
		db:  db,
		log: log,
	}
}

func (r *cart) Create(ctx context.Context, cart *models.Cart) error {
	if err := r.db.QueryRowContext(ctx, CreateCart,
		cart.UserID,
	).Scan(&cart.ID, &cart.CreatedAt, &cart.UpdatedAt); err != nil {
		r.log.Error("could not create cart:" + err.Error())
		return err
	}

	return nil
}

func (r *cart) GetByID(ctx context.Context, id string) (*models.Cart, error) {
	cart := models.Cart{}
	if err := r.db.QueryRowContext(ctx, GetCart, id).Scan(
		&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt,
		&cart.DeletedAt); err != nil {
		r.log.Error("could not get cart by id: ", err.Error())
		return nil, err
	}

	return &cart, nil
}

func (r *cart) AddProduct(ctx context.Context, cartProduct *models.CartProduct) error {
	if _, err := r.db.ExecContext(ctx, CreateCartProduct,
		cartProduct.CartID, cartProduct.ProductID, cartProduct.Quantity,
	); err != nil {
		r.log.Error("could not create cart product:" + err.Error())
		return err
	}

	return nil
}
func (r *cart) CheckStatus(ctx context.Context, id string) (bool, error) {
	var exists bool
	if err := r.db.QueryRowContext(ctx, CheckStatus, id).Scan(&exists); err != nil {
		r.log.Error("could not check field: " + err.Error())
		return false, err
	}

	return exists, nil
}

func (r *cart) UpdateStatus(ctx context.Context, id string, status models.CartStatus) error {
	if _, err := r.db.ExecContext(ctx, UpdateStatus, status, id); err != nil {
		r.log.Error("could not update cart status:" + err.Error())
		return err
	}

	return nil
}

func (r *cart) UpdateQuantity(ctx context.Context, cart *models.CartProduct) error {
	if _, err := r.db.ExecContext(ctx, UpdateQuantity, cart.Quantity, cart.CartID, cart.ProductID); err != nil {
		r.log.Error("could not update cart quantity:" + err.Error())
		return err
	}

	return nil
}
