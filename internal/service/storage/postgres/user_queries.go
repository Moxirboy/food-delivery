package queries

const (
	createUser = `
		INSERT INTO users
		(first_name,
		 last_name,
		 email,
		 position,
		 password
		 )
		VALUES ($1, $2, $3, $4, $5)
		returning id
`
)
