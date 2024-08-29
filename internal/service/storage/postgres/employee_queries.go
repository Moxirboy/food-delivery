package queries

const (
	checkFieldEmployee = `
		SELECT count(1)
		FROM employee_details
		WHERE %s = $1 AND deleted_at IS NULL
	`

	createEmployee = `
		INSERT INTO employee_details
		(first_name,
		 last_name,
		 email,
		 position,
		 password
		 )
		VALUES ($1, $2, $3, $4, $5)
		returning id
	`

	getEmployeeByAuthCred = `
		SELECT  id,
				first_name,
				last_name,
				position,
		FROM employee_details
		WHERE deleted_at IS NULL
		AND email = $1
		AND password = $2
	`
)
