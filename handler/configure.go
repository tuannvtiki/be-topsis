package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/nguyenvantuan2391996/be-topsis/internal/domain/usecase"
)

type Handler struct {
	userDomain        *usecase.UserDomain
	standardDomain    *usecase.StandardDomain
	scoreRatingDomain *usecase.ScoreRatingDomain
	consultDomain     *usecase.ConsultDomain
	validate          *validator.Validate
}

func NewHandler(
	userDomain *usecase.UserDomain,
	standardDomain *usecase.StandardDomain,
	scoreRatingDomain *usecase.ScoreRatingDomain,
	consultDomain *usecase.ConsultDomain,
	validate *validator.Validate) *Handler {
	return &Handler{
		userDomain:        userDomain,
		standardDomain:    standardDomain,
		scoreRatingDomain: scoreRatingDomain,
		consultDomain:     consultDomain,
		validate:          validate,
	}
}
