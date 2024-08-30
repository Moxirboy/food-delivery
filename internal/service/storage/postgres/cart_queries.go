package postgres

const (
	CreateCart = `
	insert into carts (user_id) values ($1) returning id, created_at, updated_at;
`
	GetCart = `
	select
		id, user_id, created_at, updated_at, deleted_at
	from carts
	where user_id = $1;
`
	CreateCartProduct = `
	insert into cart_products(cart_id, product_id, quantity) values ($1, $2, $3);
`
	GetCartProductsList = `
	select
		cart_id, product_id, quantity
	from cart_products
	where cart_id = $1;

	`
	CheckStatus = `
	select exists(select 1 from carts where id = $1 and status = 'PENDING' and deleted_at is null);`
	UpdateStatus = `
	update carts set status = $1 where id = $2;
`

	UpdateQuantity = `
	update cart_products set quantity = $1 where cart_id = $2 and product_id = $3;
`
)
