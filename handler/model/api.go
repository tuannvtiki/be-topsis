package model

const SUCCESS = "Success"

type DeletedResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}
