package customer

import (
	"errors"

	"github.com/EduardoMark/BillingCore/pkg/validate"
	"github.com/go-playground/validator/v10"
)

type CreateCustomerPayload struct {
	ID               string `json:"id"`
	AccountID        string `json:"account_id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	CpfCnpj          string `json:"cpf_cnpj"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	AddressNumber    string `json:"address_number"`
	Province         string `json:"province"`
	PostalCode       string `json:"postal_code"`
	ExternalID       string `json:"external_id"`
	ExternalPlatform string `json:"external_platform"`
}

type UpdateCustomerPayload struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	CpfCnpj          string `json:"cpf_cnpj"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	AddressNumber    string `json:"address_number"`
	Province         string `json:"province"`
	PostalCode       string `json:"postal_code"`
	ExternalID       string `json:"external_id"`
	ExternalPlatform string `json:"external_platform"`
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
