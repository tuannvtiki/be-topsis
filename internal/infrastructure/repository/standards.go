package repository

import "gorm.io/gorm"

type StandardRepository struct {
	db *gorm.DB
}

func NewStandardRepository(db *gorm.DB) *StandardRepository {
	return &StandardRepository{db: db}
}
