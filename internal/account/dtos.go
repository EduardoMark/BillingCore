package account

import (
	"errors"

	"github.com/EduardoMark/BillingCore/pkg/validate"
	"github.com/go-playground/validator/v10"
)

type CreateAccountPayload struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func Validate[T any](t T) []string {
	err := validate.Validate.Struct(t)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		errs := make([]string, 0, len(validationErrors))

		for _, fieldErr := range validationErrors {
			errs = append(errs, validate.FormatValidationError(fieldErr))
		}

		return errs
	}

	return nil
}
