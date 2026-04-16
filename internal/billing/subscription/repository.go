package subscription

import (
	"context"
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
	ErrSubscriptionNotFound = fmt.Errorf("subscription not found")
)

func (r *Repository) Create(ctx context.Context, subscription *Subscription) error {
	err := r.db.WithContext(ctx).Create(subscription).Error
	if err != nil {
		return fmt.Errorf("Repository.Create error: %w", err)
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Subscription, error) {
	var sub Subscription
	err := r.db.WithContext(ctx).First(&sub, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrSubscriptionNotFound
		}
		return nil, fmt.Errorf("Repository.GetByID error: %w", err)
	}

	return &sub, nil
}

func (r *Repository) GetByCustomerIDAndPlanID(ctx context.Context, customerID, planID string) (*Subscription, error) {
	var sub Subscription
	err := r.db.WithContext(ctx).First(&sub, "customer_id = ? AND plan_id = ?", customerID, planID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrSubscriptionNotFound
		}
		return nil, fmt.Errorf("Repository.GetByCustomerIDAndPlanID error: %w", err)
	}

	return &sub, nil
}

func (r *Repository) Update(ctx context.Context, sub *Subscription) error {
	err := r.db.WithContext(ctx).Updates(&sub).Error
	if err != nil {
		return fmt.Errorf("Repository.Update error: %w", err)
	}

	return nil
}
