package repository

import (
	"context"

	"topsis/internal/domain/model"
)

type ScoreRatingRepositoryInterface interface {
	BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error
	GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}) ([]*model.ScoreRating, error)
	DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error
}
