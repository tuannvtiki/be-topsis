package model

import "time"

type ScoreRating struct {
	Id        string     `json:"id"`
	UserId    string     `json:"user_id"`
	Metadata  string     `json:"metadata"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
