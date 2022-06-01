package handler

import "topsis/internal/domain/usecase"

type Handler struct {
	userDomain        *usecase.UserDomain
	standardDomain    *usecase.StandardDomain
	scoreRatingDomain *usecase.ScoreRatingDomain
}

func NewHandler(
	userDomain *usecase.UserDomain,
	standardDomain *usecase.StandardDomain,
	scoreRatingDomain *usecase.ScoreRatingDomain) *Handler {
	return &Handler{
		userDomain:        userDomain,
		standardDomain:    standardDomain,
		scoreRatingDomain: scoreRatingDomain,
	}
}
