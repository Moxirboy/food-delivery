package models

import (
	"database/sql"
)

type Product struct {
	ID          string
	Name        sql.NullString
	Description sql.NullString
	Price       sql.NullString
	Image       sql.NullString

	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
