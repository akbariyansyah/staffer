package helper

import (
	"github.com/go-playground/validator/v10"

)
func ValidateRequest(input interface{}) error {
	validate := validator.New()
	return validate.Struct(input)
}
