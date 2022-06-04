package model

type ConsultResponse struct {
	Name       string  `json:"name"`
	Similarity float32 `json:"similarity"`
}
