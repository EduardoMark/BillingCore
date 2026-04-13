package plans

import (
	"errors"

	"github.com/EduardoMark/BillingCore/pkg/validate"
	"github.com/go-playground/validator/v10"
)

type CreatePlanPayload struct {
	Name         string       `json:"name" validate:"required,min=3,max=100"`
	Description  string       `json:"description" validate:"required,min=10,max=500"`
	Price        int64        `json:"price" validate:"gte=0"`
	BillingCycle BillingCycle `json:"billing_cycle" validate:"required,oneof=monthly yearly"`
}

type UpdatePlanPayload struct {
	Name         string       `json:"name" validate:"min=3,max=100"`
	Description  string       `json:"description" validate:"min=10,max=500"`
	Price        int64        `json:"price" validate:"gte=0"`
	BillingCycle BillingCycle `json:"billing_cycle" validate:"oneof=monthly yearly"`
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
