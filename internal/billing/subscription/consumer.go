package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/EduardoMark/BillingCore/internal/integration/asaas"
)

type Consumer struct {
	service *Service
}

func NewConsumer(service *Service) *Consumer {
	return &Consumer{service: service}
}

func (c *Consumer) Handle(ctx context.Context, msg []byte) error {
	var sub Subscription
	err := json.Unmarshal(msg, &sub)
	if err != nil {
		return fmt.Errorf("Consumer.Handle unmarshal error: %w", err)
	}

	asaasApiKey := os.Getenv("ASAAS_API_KEY")
	asaasClient := asaas.NewClient(asaasApiKey)

	asaasSubscription, err := asaasClient.CreateSubscription(ctx, asaas.SubscriptionRequest{
		Customer:          sub.ExternalCustomerID,
		BillingType:       string(sub.BillingType),
		NextDueDate:       sub.NextDueDate,
		Value:             sub.Value,
		Cycle:             string(sub.Cycle),
		ExternalReference: sub.ID,
	})
	if err != nil {
		return fmt.Errorf("Consumer.Handle CreateSubscription error: %w", err)
	}

	_, err = c.service.Activate(ctx, sub.ID, asaasSubscription.ID)
	if err != nil {
		return fmt.Errorf("Consumer.Handle Activate error: %w", err)
	}

	return nil
}
