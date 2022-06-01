package repository

import (
	"context"

	"topsis/internal/domain/model"
)

type ScoreRatingRepositoryInterface interface {
	CreateScoreRating(ctx context.Context, standard *model.ScoreRating) (*model.ScoreRating, error)
	DeleteScoreRatingByQueries(ctx context.Context, standard map[string]interface{}) (*model.ScoreRating, error)
}
