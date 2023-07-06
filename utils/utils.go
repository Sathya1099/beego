package utils

import (
	"log"

	"github.com/Sathya1099/beego/models"

	"github.com/beego/beego/validation"
)

func Validate(object *models.Object) (validateErrors []string) {
	valid := validation.Validation{}
	b, err := valid.Valid(&object)
	if err != nil {
		validateErrors = append(validateErrors, err.Error())
		return validateErrors
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			validateErrors = append(validateErrors, err.Message)
		}
	}
	return validateErrors
}
