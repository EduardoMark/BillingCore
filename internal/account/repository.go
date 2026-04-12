package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrAccountNotFound    = errors.New("account not found")
)

func (r *Repository) Create(ctx context.Context, account *Account) error {
	err := r.db.WithContext(ctx).Create(&account).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return ErrEmailAlreadyExists
			}
		}

		zap.L().Error("Failed to create accounnt", zap.Error(err))
		return fmt.Errorf("Repository.Create error: %w", err)
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Account, error) {
	var account Account
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAccountNotFound
		}

		zap.L().Error("Failed to get account by ID", zap.Error(err))
		return nil, fmt.Errorf("Repository.GetByID error: %w", err)
	}

	return &account, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*Account, error) {
	var account Account
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAccountNotFound
		}

		zap.L().Error("Failed to get account by email", zap.Error(err))
		return nil, fmt.Errorf("Repository.GetByEmail error: %w", err)
	}

	return &account, nil
}
