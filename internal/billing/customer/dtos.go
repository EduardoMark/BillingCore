package customer

import (
	"errors"

	"github.com/EduardoMark/BillingCore/pkg/validate"
	"github.com/go-playground/validator/v10"
)

type CreateCustomerPayload struct {
	Name             string `json:"name" validate:"required,min=3,max=100"`
	Email            string `json:"email" validate:"required,email"`
	CpfCnpj          string `json:"cpf_cnpj" validate:"required,len=11|len=14"`
	Phone            string `json:"phone" validate:"required,len=11"`
	Address          string `json:"address" validate:"required"`
	AddressNumber    string `json:"address_number" validate:"required"`
	Province         string `json:"province" validate:"required"`
	PostalCode       string `json:"postal_code" validate:"required,len=8"`
	ExternalPlatform string `json:"external_platform" validate:"required,oneof=asaas"`
}

type UpdateCustomerPayload struct {
	Name             string `json:"name" validate:"required,min=3,max=100"`
	Email            string `json:"email" validate:"required,email"`
	CpfCnpj          string `json:"cpf_cnpj" validate:"required,len=11|len=14"`
	Phone            string `json:"phone" validate:"required,len=11"`
	Address          string `json:"address" validate:"required"`
	AddressNumber    string `json:"address_number" validate:"required"`
	Province         string `json:"province" validate:"required"`
	PostalCode       string `json:"postal_code" validate:"required,len=8"`
	ExternalPlatform string `json:"external_platform" validate:"required,oneof=asaas"`
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
