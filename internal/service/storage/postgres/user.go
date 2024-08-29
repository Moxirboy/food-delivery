package postgres

import (
	"context"
	"database/sql"
	"food-delivery/pkg/logger"
)

type userRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewUserRepository(db *sql.DB, log logger.Logger) repository.IUserRepository {
	return &userRepository{db: db, log: log}
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	var (
		UserID string
	)
	row:=u.db.QueryRow(
		ctx,
		createUser
		)
}
