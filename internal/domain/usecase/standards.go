package usecase

import "topsis/internal/domain/repository"

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
