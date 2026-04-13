package customer

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Platform string

const (
	PlatformAsaas Platform = "asaas"
)

type Customer struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	AccountID        string    `json:"account_id" gorm:"not null;index"`
	Name             string    `json:"name" gorm:"not null"`
	Email            string    `json:"email" gorm:"not null"`
	CpfCnpj          string    `json:"cpf_cnpj" gorm:"not null"`
	Phone            string    `json:"phone" gorm:"not null"`
	Address          string    `json:"address" gorm:"not null"`
	AddressNumber    string    `json:"address_number" gorm:"not null"`
	Province         string    `json:"province" gorm:"not null"`
	PostalCode       string    `json:"postal_code" gorm:"not null"`
	ExternalID       string    `json:"external_id" gorm:"index"`
	ExternalPlatform Platform  `json:"external_platform" gorm:"not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	id := fmt.Sprintf("cust_%s", uuid.New().String())
	c.ID = id
	return nil
}
