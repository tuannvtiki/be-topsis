package repository

import (
	"github.com/nguyenvantuan2391996/be-topsis/internal/domain/model"
)

//go:generate mockgen -package=repository -destination=istatisticals_mock.go -source=istatisticals.go

type IStatisticalRepositoryInterface interface {
	Create(record *model.Statistical) error
	UpdateWithMap(record *model.Statistical, queries map[string]interface{}) error
	GetByQueries(queries map[string]interface{}) (*model.Statistical, error)
}
