package queries

const (
	CreateUser = `
	Insert into users (username, password) values ($1, $2) returning id;
`
)
