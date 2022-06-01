package repository

import (
	"context"

	"gorm.io/gorm"
	"topsis/internal/domain/model"
)

type ScoreRatingRepository struct {
	db *gorm.DB
}

func NewScoreRatingRepository(db *gorm.DB) *ScoreRatingRepository {
	return &ScoreRatingRepository{db: db}
}

func (sr *ScoreRatingRepository) CreateScoreRating(ctx context.Context, standard *model.ScoreRating) (*model.ScoreRating, error) {
	return nil, nil
}
func (sr *ScoreRatingRepository) DeleteScoreRatingByQueries(ctx context.Context, standard map[string]interface{}) (*model.ScoreRating, error) {
	return nil, nil
}
