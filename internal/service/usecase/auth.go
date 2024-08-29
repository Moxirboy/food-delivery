package usecase

import (
	"database/sql"
	"food-delivery/pkg/logger"
)

type AuthUsecase struct {
	db  *sql.DB
	log logger.Logger
}

func NewAuthUsecase(
	db *sql.DB,
	log logger.Logger,
) AuthUsecase {
	return AuthUsecase{
		db:  db,
		log: log,
	}
}
