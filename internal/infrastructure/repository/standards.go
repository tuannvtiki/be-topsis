package repository

import (
	"context"

	"gorm.io/gorm"
	"topsis/internal/domain/model"
)

type StandardRepository struct {
	db *gorm.DB
}

func NewStandardRepository(db *gorm.DB) *StandardRepository {
	return &StandardRepository{db: db}
}

func (s *StandardRepository) CreateStandard(ctx context.Context, standard *model.Standard) (*model.Standard, error) {
	return nil, nil
}

func (s *StandardRepository) GetStandardByQueries(ctx context.Context, queries map[string]interface{}) (*model.Standard, error) {
	return nil, nil
}
func (s *StandardRepository) DeleteStandardByQueries(ctx context.Context, standard map[string]interface{}) (*model.Standard, error) {
	return nil, nil
}
