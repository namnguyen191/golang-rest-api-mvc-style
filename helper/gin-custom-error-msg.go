package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", e.Field(), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", e.Field(), e.Param())
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", e.Field(), e.Param())
	}
	return fmt.Sprintf("%s is not valid", e.Field())
}

func ValidationErrorsToStringArray(e validator.ValidationErrors) []string {
	var errs []string
	for _, fieldErr := range e {
		errs = append(errs, ValidationErrorToText(fieldErr))
	}

	return errs
}
