package repository

import (
	"context"

	"topsis/internal/domain/model"
)

type ScoreRatingRepositoryInterface interface {
	BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error
	GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}, sort []string) ([]*model.ScoreRating, error)
	UpdateScoreRatingWithMap(ctx context.Context, scoreRating *model.ScoreRating, data map[string]interface{}) error
	DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error
}
