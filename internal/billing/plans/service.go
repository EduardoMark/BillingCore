package plans

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

var (
	ErrInvalidPrice = errors.New("price must be greater than zero")
)

func (s *Service) Create(ctx context.Context, accountID string, payload *CreatePlanPayload) error {
	if payload.Price < 0 {
		return ErrInvalidPrice
	}

	plan := &Plan{
		AccountID:    accountID,
		Name:         payload.Name,
		Description:  payload.Description,
		Price:        payload.Price,
		BillingCycle: payload.BillingCycle,
	}

	err := s.repo.Create(ctx, plan)
	if err != nil {
		return err
	}

	return nil
}
