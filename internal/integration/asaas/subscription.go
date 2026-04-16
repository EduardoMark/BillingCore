package asaas

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type SubscriptionRequest struct {
	Customer          string   `json:"customer"`
	BillingType       string   `json:"billingType"`
	Value             float64  `json:"value"`
	NextDueDate       string   `json:"nextDueDate"`
	Discount          Discount `json:"discount"`
	Interest          Interest `json:"interest"`
	Fine              Fine     `json:"fine"`
	EndDate           string   `json:"endDate"`
	Cycle             string   `json:"cycle"`
	Description       string   `json:"description"`
	MaxPayments       int      `json:"maxPayments"`
	ExternalReference string   `json:"externalReference"`
	Split             []Split  `json:"split"`
}

type SubscriptionResponse struct {
	Object            string    `json:"object"`
	ID                string    `json:"id"`
	DateCreated       string    `json:"dateCreated"`
	Customer          string    `json:"customer"`
	PaymentLink       *string   `json:"paymentLink"`
	BillingType       string    `json:"billingType"`
	Cycle             string    `json:"cycle"`
	Value             float64   `json:"value"`
	NextDueDate       string    `json:"nextDueDate"`
	EndDate           string    `json:"endDate"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Discount          *Discount `json:"discount"`
	Fine              *Fine     `json:"fine"`
	Interest          *Interest `json:"interest"`
	Deleted           bool      `json:"deleted"`
	MaxPayments       int       `json:"maxPayments"`
	ExternalReference *string   `json:"externalReference"`
	CheckoutSession   string    `json:"checkoutSession"`
	Split             []Split   `json:"split"`
}

type Discount struct {
	Value            float64 `json:"value"`
	DueDateLimitDays int     `json:"dueDateLimitDays"`
	Type             string  `json:"type"`
}

type Fine struct {
	Value float64 `json:"value"`
}

type Interest struct {
	Value float64 `json:"value"`
}

type Split struct {
	WalletID          string   `json:"walletId"`
	FixedValue        float64  `json:"fixedValue"`
	PercentualValue   *float64 `json:"percentualValue"`
	ExternalReference *string  `json:"externalReference"`
	Description       *string  `json:"description"`
	Status            string   `json:"status"`
	DisabledReason    *string  `json:"disabledReason"`
}

func (c *Client) CreateSubscription(ctx context.Context, payload SubscriptionRequest) (*SubscriptionResponse, error) {
	respBody, err := c.DoRequest(ctx, http.MethodPost, "/subscriptions", payload)
	if err != nil {
		zap.L().Error("Asaas.CreateSubscription error", zap.Error(err))
		return nil, fmt.Errorf("Asaas.CreateSubscription error: %w", err)
	}

	resp, err := DecodeResponse[SubscriptionResponse](respBody)
	if err != nil {
		zap.L().Error("Asaas.CreateSubscription error", zap.Error(err))
		return nil, fmt.Errorf("Asaas.CreateSubscription error: %w", err)
	}

	return resp, nil
}
