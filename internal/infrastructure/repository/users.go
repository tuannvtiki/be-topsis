package repository

import (
	"context"

	"gorm.io/gorm"
	"topsis/internal/domain/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := user.Model.GenerateID(); err != nil {
		return nil, err
	}
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
