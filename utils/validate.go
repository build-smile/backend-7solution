package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func ValidateRequest(req interface{}) error {
	err := Validate.Struct(req)
	errMsg := ""
	if err != nil {
		// Return validation errors
		for _, err := range err.(validator.ValidationErrors) {
			errMsg = errMsg + fmt.Sprintf("{Field: %s Error: %s} ", err.Field(), err.Tag())
		}
		return NewCustomError(http.StatusBadRequest, errMsg)
	}
	return nil
}
