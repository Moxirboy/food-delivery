package postgres

import (
	"context"
	"database/sql"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
)

type order struct {
	db  *sql.DB
	log logger.Logger
}

func NewOrder(db *sql.DB, log logger.Logger) repo.IOrderRepository {
	return &order{
		db:  db,
		log: log,
	}
}

func (o *order) CreateOrder(ctx context.Context, order models.Order) error {

	_, err := o.db.ExecContext(ctx, "INSERT INTO orders(cart_id, status) VALUES($1, $2)", order.CartID, order.Status)
	if err != nil {
		o.log.Errorf("Error while creating order: %v", err)
		return err
	}
	return nil
}

func (o *order) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	order := models.Order{}
	if err := o.db.QueryRowContext(ctx, "SELECT id, cart_id, status FROM orders WHERE id = $1", id).Scan(&order.ID, &order.CartID, &order.Status); err != nil {
		o.log.Errorf("Error while getting order by id: %v", err)
		return nil, err
	}
	return &order, nil
}

func (o *order) GetOrders(ctx context.Context) ([]*models.Order, error) {
	rows, err := o.db.QueryContext(ctx, "SELECT id, cart_id, status FROM orders")
	if err != nil {
		o.log.Errorf("Error while getting orders: %v", err)
		return nil, err
	}
	defer rows.Close()

	orders := make([]*models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		if err := rows.Scan(&order.ID, &order.CartID, &order.Status); err != nil {
			o.log.Errorf("Error while scanning orders: %v", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (o *order) UpdateOrder(ctx context.Context, order models.Order) error {
	_, err := o.db.ExecContext(ctx, "UPDATE orders SET cart_id = $1, status = $2 WHERE id = $3", order.CartID, order.Status, order.ID)
	if err != nil {
		o.log.Errorf("Error while updating order: %v", err)
		return err
	}
	return nil
}

func (o *order) DeleteOrder(ctx context.Context, id string) error {
	_, err := o.db.ExecContext(ctx, "DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		o.log.Errorf("Error while deleting order: %v", err)
		return err
	}
	return nil
}

func (o *order) GetOrderByCartID(ctx context.Context, cartID int) (*models.Order, error) {
	order := models.Order{}
	if err := o.db.QueryRowContext(ctx, "SELECT id, cart_id, status FROM orders WHERE cart_id = $1", cartID).Scan(&order.ID, &order.CartID, &order.Status); err != nil {
		o.log.Errorf("Error while getting order by cart id: %v", err)
		return nil, err
	}
	return &order, nil
}

func (o *order) GetOrdersByStatus(ctx context.Context, status string) ([]*models.Order, error) {
	rows, err := o.db.QueryContext(ctx, "SELECT id, cart_id, status FROM orders WHERE status = $1", status)
	if err != nil {
		o.log.Errorf("Error while getting orders by status: %v", err)
		return nil, err
	}
	defer rows.Close()

	orders := make([]*models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		if err := rows.Scan(&order.ID, &order.CartID, &order.Status); err != nil {
			o.log.Errorf("Error while scanning orders by status: %v", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}



func (o *order) GetOrdersByCourierID(ctx context.Context, userID int) ([]*models.Order, error) {
	rows, err := o.db.QueryContext(ctx, "SELECT o.id, o.cart_id, o.status FROM orders o JOIN carts c ON o.cart_id = c.id WHERE c.curier_id = $1 and status= delivering ", userID)
	if err != nil {
		o.log.Errorf("Error while getting orders by user id: %v", err)
		return nil, err
	}
	defer rows.Close()

	orders := make([]*models.Order, 0)
	for rows.Next() {
		order := models.Order{}
		if err := rows.Scan(&order.ID, &order.CartID, &order.Status); err != nil {
			o.log.Errorf("Error while scanning orders by user id: %v", err)
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (o *order) GetProductOrderDetails(ctx context.Context, id string) (*models.OrderProducts, error) {
	rows,err:=o.db.QueryContext(ctx,"SELECT o.id, o.cart_id, o.status, p.id, p.name, p.price, p.image FROM orders o JOIN carts c ON o.cart_id = c.id JOIN cart_products cp ON c.id = cp.cart_id JOIN products p ON cp.product_id = p.id WHERE o.id = $1 ",id)
	if err != nil {
		o.log.Errorf("Error while getting product order details: %v", err)
		return	nil, err

	}
	defer rows.Close()
	order:=models.OrderProducts{}
	order.Product=make([]models.Product,0)
	for rows.Next() {
		product:=models.Product{}
		if err := rows.Scan(&order.ID, &order.OrderID, &order.Status, &product.ID, &product.Name, &product.Price, &product.Image); err != nil {
			o.log.Errorf("Error while scanning product order details: %v", err)
			return nil, err
		}
		order.Product=append(order.Product,product)
	}		
	return &order,nil
}