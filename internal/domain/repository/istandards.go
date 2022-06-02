package repository

import (
	"context"

	"topsis/internal/domain/model"
)

type StandardRepositoryInterface interface {
	CreateStandard(ctx context.Context, standard *model.Standard) (*model.Standard, error)
	GetStandardByQueries(ctx context.Context, queries map[string]interface{}) ([]*model.Standard, error)
	DeleteStandardByQueries(ctx context.Context, queries map[string]interface{}) error
}
