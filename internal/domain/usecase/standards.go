package usecase

import (
	"context"

	"topsis/handler/model"
	modelDomain "topsis/internal/domain/model"
	"topsis/internal/domain/repository"
)

type StandardDomain struct {
	standardRepo repository.StandardRepositoryInterface
}

func NewStandardDomain(
	standardRepo repository.StandardRepositoryInterface,
) *StandardDomain {
	return &StandardDomain{
		standardRepo: standardRepo,
	}
}

func (s *StandardDomain) CreateStandard(ctx context.Context, standardRequest *model.StandardRequest) (*modelDomain.Standard, error) {
	return s.standardRepo.CreateStandard(ctx, &modelDomain.Standard{
		UserId:       standardRequest.UserID,
		StandardName: standardRequest.StandardName,
		Weight:       standardRequest.Weight,
	})
}

func (s *StandardDomain) GetStandards(ctx context.Context, standard *model.Standard) ([]*modelDomain.Standard, error) {
	return s.standardRepo.GetStandardByQueries(ctx, map[string]interface{}{
		"user_id": standard.UserID,
	})
}

func (s *StandardDomain) DeleteStandard(ctx context.Context, standard *model.Standard) error {
	return s.standardRepo.DeleteStandardByQueries(ctx, map[string]interface{}{
		"id": standard.ID,
	})
}
