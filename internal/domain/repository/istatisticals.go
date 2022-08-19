package repository

import (
	"topsis/internal/domain/model"
)

type StatisticalRepositoryInterface interface {
	Create(record *model.Statistical) error
	UpdateWithMap(record *model.Statistical, queries map[string]interface{}) error
	GetByQueries(queries map[string]interface{}) (*model.Statistical, error)
}
