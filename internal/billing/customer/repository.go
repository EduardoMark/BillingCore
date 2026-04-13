package customer

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

var (
	ErrCustomerNotFound = fmt.Errorf("customer not found")
)

func (r *Repository) Create(ctx context.Context, customer *Customer) error {
	err := r.db.WithContext(ctx).Create(&customer).Error
	if err != nil {
		return fmt.Errorf("Repository.Create error: %w", err)
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Customer, error) {
	var customer Customer
	err := r.db.WithContext(ctx).First(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}

		return nil, fmt.Errorf("Repository.GetByID error: %w", err)
	}

	return &customer, nil
}

func (r *Repository) GetByExternalID(ctx context.Context, externalID string) (*Customer, error) {
	var customer Customer
	err := r.db.WithContext(ctx).Where("external_id = ?", externalID).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}

		return nil, fmt.Errorf("Repository.GetByExternalID error: %w", err)
	}

	return &customer, nil
}

func (r *Repository) GetAllByAccountID(ctx context.Context, accountID string) ([]*Customer, error) {
	var customers []*Customer
	err := r.db.WithContext(ctx).Where("account_id = ?", accountID).Find(&customers).Error
	if err != nil {
		return nil, fmt.Errorf("Repository.GetAllByAccountID error: %w", err)
	}

	return customers, nil
}

func (r *Repository) Update(ctx context.Context, customer *Customer) error {
	err := r.db.WithContext(ctx).Updates(&customer)
	if err != nil {
		return fmt.Errorf("Repository.Update error: %w", err.Error)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Customer{}, id).Error
	if err != nil {
		return fmt.Errorf("Repository.Delete error: %w", err)
	}

	return nil
}
