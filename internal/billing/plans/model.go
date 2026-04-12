package plans

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingCycle string

const (
	Monthly BillingCycle = "monthly"
	Yearly  BillingCycle = "yearly"
)

type Plan struct {
	ID           string       `json:"id" gorm:"primaryKey"`
	AccountID    string       `json:"account_id" gorm:"index;not null"`
	Name         string       `json:"name" gorm:"not null"`
	Description  string       `json:"description" gorm:"not null"`
	Price        int64        `json:"price" gorm:"not null"`
	BillingCycle BillingCycle `json:"billing_cycle" gorm:"not null"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *Plan) BeforeCreate(tx *gorm.DB) error {
	p.ID = fmt.Sprintf("plan_%s", uuid.New().String())
	return nil
}
