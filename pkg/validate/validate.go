package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func Validate[T any](t T) []string {
	err := validate.Struct(t)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		errs := make([]string, 0, len(validationErrors))

		for _, fieldErr := range validationErrors {
			errs = append(errs, FormatValidationError(fieldErr))
		}

		return errs
	}

	return nil
}

func FormatValidationError(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, fe.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of the following: %s", field, fe.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
