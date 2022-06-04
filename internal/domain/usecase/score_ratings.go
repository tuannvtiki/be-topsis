package usecase

import (
	"context"

	"topsis/handler/model"
	modelDomain "topsis/internal/domain/model"
	"topsis/internal/domain/repository"
)

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

func (s *ScoreRatingDomain) BulkCreateScoreRating(ctx context.Context, scoreRatingReq []*model.ScoreRatingRequest) error {
	var scoreRatings []*modelDomain.ScoreRating
	for _, value := range scoreRatingReq {
		v := &modelDomain.ScoreRating{
			UserId:   value.UserId,
			Metadata: value.Metadata,
		}
		if err := v.GenerateID(); err != nil {
			return err
		}
		scoreRatings = append(scoreRatings, v)
	}
	return s.scoreRatingRepo.BulkCreateScoreRating(ctx, scoreRatings)
}

func (s *ScoreRatingDomain) GetListScoreRating(ctx context.Context, scoreRating *model.ScoreRating) ([]*modelDomain.ScoreRating, error) {
	return s.scoreRatingRepo.GetScoreRatingByListQueries(ctx, map[string]interface{}{
		"user_id": scoreRating.UserId,
	})
}

func (s *ScoreRatingDomain) DeleteScoreRating(ctx context.Context, scoreRating *model.ScoreRating) error {
	return s.scoreRatingRepo.DeleteScoreRatingByQueries(ctx, map[string]interface{}{
		"id": scoreRating.ID,
	})
}
