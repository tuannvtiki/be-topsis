package model

import "topsis/handler/model"

type Standard struct {
	Model
	UserId       string `json:"user_id"`
	StandardName string `json:"standard_name"`
	Weight       int    `json:"weight"`
	Type         string `json:"type"`
}

func (s *Standard) ToResponse() *model.StandardResponse {
	return &model.StandardResponse{
		ID:           s.ID,
		UserID:       s.UserId,
		StandardName: s.StandardName,
		Weight:       s.Weight,
	}
}
