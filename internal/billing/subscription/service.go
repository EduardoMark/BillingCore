package subscription

import (
	"context"
	"errors"
	"fmt"

	"github.com/EduardoMark/BillingCore/internal/infra/rabbitmq"
)

type Service struct {
	repo     *Repository
	producer *rabbitmq.Producer
}

func NewService(repository *Repository, producer *rabbitmq.Producer) *Service {
	return &Service{repo: repository, producer: producer}
}

var (
	ErrCustomerAlreadyHasSubscription = fmt.Errorf("customer already has a subscription for this plan")
)

func (s *Service) Create(ctx context.Context, accountID string, payload *CreateSubscriptionRequest) (*Subscription, error) {
	hasSub, err := s.repo.GetByCustomerIDAndPlanID(ctx, payload.CustomerID, payload.PlanID)
	if err != nil && !errors.Is(err, ErrSubscriptionNotFound) {
		return nil, err
	}
	if hasSub != nil {
		return nil, ErrCustomerAlreadyHasSubscription
	}

	sub := &Subscription{
		AccountID:          accountID,
		PlanID:             payload.PlanID,
		CustomerID:         payload.CustomerID,
		ExternalCustomerID: payload.ExternalCustomerID,
		Status:             StatusPending,
		BillingType:        BillingType(payload.BillingType),
		Value:              payload.Value,
		NextDueDate:        payload.NextDueDate,
		Cycle:              Cycle(payload.Cycle),
	}

	err = s.repo.Create(ctx, sub)
	if err != nil {
		return nil, err
	}

	err = s.producer.Publish(ctx, "subscription.created", sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *Service) Activate(ctx context.Context, id, externalSubscriptionID string) (*Subscription, error) {
	sub, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	sub.Status = StatusActive
	sub.ExternalSubscriptionID = externalSubscriptionID

	err = s.repo.Update(ctx, sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
