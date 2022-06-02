package handler

import (
	"github.com/go-playground/validator/v10"
	"topsis/internal/domain/usecase"
)

type Handler struct {
	userDomain        *usecase.UserDomain
	standardDomain    *usecase.StandardDomain
	scoreRatingDomain *usecase.ScoreRatingDomain
	validate          *validator.Validate
}

func NewHandler(
	userDomain *usecase.UserDomain,
	standardDomain *usecase.StandardDomain,
	scoreRatingDomain *usecase.ScoreRatingDomain,
	validate *validator.Validate) *Handler {
	return &Handler{
		userDomain:        userDomain,
		standardDomain:    standardDomain,
		scoreRatingDomain: scoreRatingDomain,
		validate:          validate,
	}
}
