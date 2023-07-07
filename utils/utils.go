package utils

import (
	"github.com/Sathya1099/beego/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/validation"
)

func Validate(object *models.Object) (validateErrors []string) {
	valid := validation.Validation{}
	b, err := valid.Valid(object)
	if err != nil {
		logs.Error(err)
	}
	if !b {
		for _, err := range valid.Errors {
			logs.Error(err)
			validateErrors = append(validateErrors, err.Message)
		}
	}
	return validateErrors
}

func ValidateForUpdate(object *models.Object) (validateErrors []string) {
	valid := validation.Validation{}

	if object.Score != 0 {
		valid.Range(object.Score, 0, 500, "score")
	}
	if object.PlayerName != "" {
		valid.MinSize(object.PlayerName, 3, "player_name")
	}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err)
			validateErrors = append(validateErrors, err.Message)
		}
	}
	return validateErrors
}
