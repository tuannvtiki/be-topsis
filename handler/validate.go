package handler

import (
	"errors"

	"github.com/nguyenvantuan2391996/be-topsis/handler/constant"
	"github.com/nguyenvantuan2391996/be-topsis/handler/model"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) ValidateRequest(requestStruct any) []*model.ErrorValidateRequest {
	var (
		ve      validator.ValidationErrors
		listErr []*model.ErrorValidateRequest
	)

	err := h.validate.Struct(requestStruct)
	if errors.As(err, &ve) {
		for _, e := range ve {
			listErr = append(listErr, &model.ErrorValidateRequest{
				Field:   e.Field(),
				Message: constant.MapErrorValidateRequest[e.Field()],
			})
		}
	}

	return listErr
}
