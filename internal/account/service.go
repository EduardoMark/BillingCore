package account

import (
	"fmt"

	"github.com/EduardoMark/BillingCore/pkg/hashing"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(payload *CreateAccountPayload) (*Account, error) {
	passwordHash, err := hashing.HashPassword(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("Service.Create error: %w", err)
	}

	acc := &Account{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: passwordHash,
	}

	err = s.repo.Create(acc)
	if err != nil {
		return nil, fmt.Errorf("Service.Create error: %w", err)
	}

	return acc, nil
}

func (s *Service) GetByID(id string) (*Account, error) {
	acc, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("Service.GetByID error: %w", err)
	}

	return acc, nil
}

func (s *Service) GetByEmail(email string) (*Account, error) {
	acc, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Service.GetByEmail error: %w", err)
	}

	return acc, nil
}
