package repository

import "gorm.io/gorm"

type ScoreRatingRepository struct {
	db *gorm.DB
}

func NewScoreRatingRepository(db *gorm.DB) *ScoreRatingRepository {
	return &ScoreRatingRepository{db: db}
}
