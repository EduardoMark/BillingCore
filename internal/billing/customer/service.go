package customer

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, accountID string, payload *CreateCustomerPayload) (*Customer, error) {
	customer := &Customer{
		AccountID:        accountID,
		Name:             payload.Name,
		Email:            payload.Email,
		CpfCnpj:          payload.CpfCnpj,
		Phone:            payload.Phone,
		Address:          payload.Address,
		AddressNumber:    payload.AddressNumber,
		Province:         payload.Province,
		PostalCode:       payload.PostalCode,
		ExternalID:       payload.ExternalID,
		ExternalPlatform: Platform(payload.ExternalPlatform),
	}

	err := s.repo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Customer, error) {
	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) GetByExternalID(ctx context.Context, externalID string) (*Customer, error) {
	customer, err := s.repo.GetByExternalID(ctx, externalID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) GetAllByAccountID(ctx context.Context, accountID string) ([]*Customer, error) {
	customers, err := s.repo.GetAllByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *Service) Update(ctx context.Context, id string, payload UpdateCustomerPayload) (*Customer, error) {
	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	customer.Name = payload.Name
	customer.Email = payload.Email
	customer.CpfCnpj = payload.CpfCnpj
	customer.Phone = payload.Phone
	customer.Address = payload.Address
	customer.AddressNumber = payload.AddressNumber
	customer.Province = payload.Province
	customer.PostalCode = payload.PostalCode
	customer.ExternalID = payload.ExternalID
	customer.ExternalPlatform = Platform(payload.ExternalPlatform)

	err = s.repo.Update(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
