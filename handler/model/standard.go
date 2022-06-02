package model

type StandardRequest struct {
	UserID       string `json:"user_id" validate:"required"`
	StandardName string `json:"standard_name" validate:"required,gte=3,lte=30"`
	Weight       int    `json:"weight" validate:"required"`
}

type StandardResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	StandardName string `json:"standard_name"`
	Weight       int    `json:"weight"`
}

type Standard struct {
	ID     string
	UserID string
}
