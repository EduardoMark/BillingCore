package asaas

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type CreateCustomerRequest struct {
	Name                 string `json:"name"`
	CpfCnpj              string `json:"cpfCnpj"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	MobilePhone          string `json:"mobilePhone"`
	Address              string `json:"address"`
	AddressNumber        string `json:"addressNumber"`
	Complement           string `json:"complement"`
	Province             string `json:"province"`
	PostalCode           string `json:"postalCode"`
	ExternalReference    string `json:"externalReference"`
	NotificationDisabled bool   `json:"notificationDisabled"`
	AdditionalEmails     string `json:"additionalEmails"`
	MunicipalInscription string `json:"municipalInscription"`
	StateInscription     string `json:"stateInscription"`
	Observations         string `json:"observations"`
	GroupName            string `json:"groupName"`
	Company              string `json:"company"`
	ForeignCustomer      bool   `json:"foreignCustomer"`
}

type CreateCustomerResponse struct {
	Object               string `json:"object"`
	ID                   string `json:"id"`
	DateCreated          string `json:"dateCreated"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	MobilePhone          string `json:"mobilePhone"`
	Address              string `json:"address"`
	AddressNumber        string `json:"addressNumber"`
	Complement           string `json:"complement"`
	Province             string `json:"province"`
	City                 int    `json:"city"`
	CityName             string `json:"cityName"`
	State                string `json:"state"`
	Country              string `json:"country"`
	PostalCode           string `json:"postalCode"`
	CpfCnpj              string `json:"cpfCnpj"`
	PersonType           string `json:"personType"`
	Deleted              bool   `json:"deleted"`
	AdditionalEmails     string `json:"additionalEmails"`
	ExternalReference    string `json:"externalReference"`
	NotificationDisabled bool   `json:"notificationDisabled"`
	Observations         string `json:"observations"`
	ForeignCustomer      bool   `json:"foreignCustomer"`
}

func (c *Client) CreateCustomer(ctx context.Context, payload *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	respBody, err := c.DoRequest(ctx, http.MethodPost, "/customers", payload)
	if err != nil {
		zap.L().Error("Asaas.CreateCustomer error", zap.Error(err))
		return nil, fmt.Errorf("Asaas.CreateCustomer error: %w", err)
	}

	var resp CreateCustomerResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		zap.L().Error("Asaas.CreateCustomer error", zap.Error(err))
		return nil, fmt.Errorf("Asaas.CreateCustomer error: %w", err)
	}

	return &resp, nil
}
