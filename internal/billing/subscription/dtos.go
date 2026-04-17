package subscription

type CreateSubscriptionRequest struct {
	PlanID             string  `json:"plan_id" validate:"required"`
	CustomerID         string  `json:"customer_id" validate:"required"`
	ExternalCustomerID string  `json:"external_customer_id" validate:"required"`
	BillingType        string  `json:"billing_type" validate:"required,oneof=BOLETO CREDIT_CARD PIX"`
	Value              float64 `json:"value" validate:"required,gte=0"`
	NextDueDate        string  `json:"next_due_date" validate:"required"`
	Cycle              string  `json:"cycle" validate:"required,oneof=MONTHLY YEARLY"`
}

type IdempotencyValue struct {
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}
