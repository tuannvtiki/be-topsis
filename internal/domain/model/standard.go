package model

import "time"

type Standard struct {
	Id           string     `json:"id"`
	UserId       string     `json:"user_id"`
	StandardName string     `json:"standard_name"`
	Weight       int        `json:"weight"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
