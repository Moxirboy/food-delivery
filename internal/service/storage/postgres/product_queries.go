package queries

const (
	CreateProduct = `
	insert into products (name, description, price, image) values ($1, $2, $3, $4) returning id, created_at, updated_at, deleted_at;
`
	GetProduct = `
	select
		id, name, description, price, image, created_at, updated_at, deleted_at
	from products
	where id = $1;
`
	GetProductList = `
	select
		id, name, description, price, image, created_at, updated_at, deleted_at
	from products
	where deleted_at is null;
`
)
