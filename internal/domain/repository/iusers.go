package repository

import (
	"context"

	"topsis/internal/domain/model"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}
