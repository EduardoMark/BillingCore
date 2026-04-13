package database

import (
	"fmt"
	"os"

	"github.com/EduardoMark/BillingCore/internal/account"
	"github.com/EduardoMark/BillingCore/internal/billing/customer"
	"github.com/EduardoMark/BillingCore/internal/billing/plans"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error on connect with db: %w", err)
	}

	DB = db

	return DB, nil
}

func Migrate() error {
	return DB.AutoMigrate(
		&account.Account{},
		&plans.Plan{},
		&customer.Customer{},
	)
}
