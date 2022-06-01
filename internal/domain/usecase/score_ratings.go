package usecase

import "topsis/internal/domain/repository"

type ScoreRatingDomain struct {
	scoreRatingRepo repository.ScoreRatingRepositoryInterface
}

func NewScoreRatingDomain(
	scoreRatingRepo repository.ScoreRatingRepositoryInterface,
) *ScoreRatingDomain {
	return &ScoreRatingDomain{
		scoreRatingRepo: scoreRatingRepo,
	}
}
