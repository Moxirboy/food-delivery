package repo

import (
	"context"
	"food-delivery/internal/models"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, login, password string) (*models.User, error)
	CheckField(
		ctx context.Context,
		field, value string,
	) (bool, error)
}
