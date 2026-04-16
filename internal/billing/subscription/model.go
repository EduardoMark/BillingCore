package subscription

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillingType string

const (
	BillingTypeBoleto     BillingType = "BOLETO"
	BillingTypeCreditCard BillingType = "CREDIT_CARD"
	BillingTypePix        BillingType = "PIX"
)

type Cycle string

const (
	CycleMonthly Cycle = "MONTHLY"
	CycleYearly  Cycle = "YEARLY"
)

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusCanceled Status = "CANCELED"
	StatusInactive Status = "INACTIVE"
	StatusPending  Status = "PENDING"
)

type Subscription struct {
	ID                     string      `json:"id" gorm:"primaryKey"`
	AccountID              string      `json:"account_id" gorm:"not null"`
	PlanID                 string      `json:"plan_id" gorm:"not null"`
	CustomerID             string      `json:"customer_id" gorm:"not null"`
	ExternalCustomerID     string      `json:"external_customer_id" gorm:"not null"`
	ExternalSubscriptionID string      `json:"external_subscription_id" gorm:"not null"`
	Status                 Status      `json:"status" gorm:"not null"`
	BillingType            BillingType `json:"billing_type" gorm:"not null"`
	Value                  float64     `json:"value" gorm:"not null"`
	NextDueDate            string      `json:"next_due_date" gorm:"not null"`
	Cycle                  Cycle       `json:"cycle" gorm:"not null"`
	CreatedAt              time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

func (s *Subscription) BeforeCreate(tx *gorm.DB) error {
	id := fmt.Sprintf("sub_%s", uuid.New().String())
	s.ID = id
	return nil
}
