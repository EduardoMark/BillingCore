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

func (s *Service) Create(ctx context.Context, accountID string, payload *CreatePlanPayload) (*Plan, error) {
	if payload.Price < 0 {
		return nil, ErrInvalidPrice
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
		return nil, err
	}

	return plan, nil
}

func (s *Service) GetOne(ctx context.Context, id string) (*Plan, error) {
	plan, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (s *Service) GetAll(ctx context.Context, accountID string) ([]*Plan, error) {
	plans, err := s.repo.ListByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (s *Service) Update(ctx context.Context, id string, payload *UpdatePlanPayload) (*Plan, error) {
	plan, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	plan.Name = payload.Name
	plan.Description = payload.Description
	plan.Price = payload.Price
	plan.BillingCycle = payload.BillingCycle

	err = s.repo.Update(ctx, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
