package constant

var (
	ParseRequestFail        = "Parse request fail"
	CreateNewUserFail       = "Create new user fail"
	CreateNewStandardFail   = "Create new standard fail"
	DeleteStandardFail      = "Delete standard fail"
	GetStandardsFail        = "Get list standard fail"
	MapErrorValidateRequest = map[string]string{
		"Name":         "User name is invalid (suggest: 3-20 characters)",
		"StandardName": "Standard name is invalid (suggest: 3-30 characters)",
		"Weight":       "Weight is required",
	}
)
