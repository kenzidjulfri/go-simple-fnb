package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

type ValidationError struct {
	FieldName   string `json:"field"`
	RuleMessage string `json:"rule"`
}

func GenerateErrorValidationResponse(err error) ErrorResponse {
	var response ErrorResponse

	for _, err := range err.(validator.ValidationErrors) {
		validationError := ValidationError{
			FieldName:   err.Field(),
			RuleMessage: fmt.Sprintf("%s is %s", err.Field(), err.Tag()),
		}
		response.Errors = append(response.Errors, validationError)
	}

	return response
}
