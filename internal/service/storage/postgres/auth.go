package postgres

import (
	"database/sql"
	"food-delivery/internal/service/storage/postgres/queries"
	"food-delivery/pkg/logger"
)

type auth struct {
	sql *sql.DB
	log logger.Logger
}

func NewAuth(sql *sql.DB, logger logger.Logger) auth {
	return auth{
		sql: sql,
		log: logger,
	}
}

func (a *auth) create() {
	_ = queries.CreateUser
}
