package plans

type CreatePlanPayload struct {
	Name         string       `json:"name" validate:"required,min=3,max=100"`
	Description  string       `json:"description" validate:"required,min=10,max=500"`
	Price        int64        `json:"price" validate:"gte=0"`
	BillingCycle BillingCycle `json:"billing_cycle" validate:"required,oneof=monthly yearly"`
}

type UpdatePlanPayload struct {
	Name         string       `json:"name" validate:"min=3,max=100"`
	Description  string       `json:"description" validate:"min=10,max=500"`
	Price        int64        `json:"price" validate:"gte=0"`
	BillingCycle BillingCycle `json:"billing_cycle" validate:"oneof=monthly yearly"`
}
