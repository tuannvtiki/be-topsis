package model

type ParamUploadImage struct {
	Path   string `json:"path"`
	ApiKey string `json:"apiKey"`
}

type ImageInfo struct {
	Image struct {
		Url string `json:"url"`
	} `json:"image"`
}
