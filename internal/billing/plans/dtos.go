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

func (p *CreatePlanPayload) Validate() []string {
	err := validate.Validate.Struct(p)
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
