package plans

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
	ErrPlanNotFound = errors.New("plan not found")
)

func (r *Repository) Create(ctx context.Context, plan *Plan) error {
	err := r.db.WithContext(ctx).Create(&plan).Error
	if err != nil {
		return fmt.Errorf("Repository.Create error: %w", err)
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Plan, error) {
	var plan Plan
	err := r.db.WithContext(ctx).First(&plan, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPlanNotFound
		}
		return nil, fmt.Errorf("Repository.GetByID error: %w", err)
	}

	return &plan, nil
}

func (r *Repository) ListByAccountID(ctx context.Context, accountID string) ([]*Plan, error) {
	var plans []*Plan
	err := r.db.WithContext(ctx).Where("account_id = ?", accountID).Find(&plans).Error
	if err != nil {
		return nil, fmt.Errorf("Repository.ListByAccountID error: %w", err)
	}

	return plans, nil
}

func (r *Repository) Update(ctx context.Context, plan *Plan) error {
	err := r.db.WithContext(ctx).Updates(&plan).Error
	if err != nil {
		return fmt.Errorf("Repository.Update error: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Plan{}, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("Repository.Delete error: %w", err)
	}

	return nil
}
