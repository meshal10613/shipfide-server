package dto

import (
	"time"

	"server/internal/models"
)

type ReceiverFraudProfileResponse struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Phone          string           `json:"phone"`
	Email          *string          `json:"email,omitempty"`
	TotalOrders    int              `json:"totalOrders"`
	TotalDelivered int              `json:"totalDelivered"`
	FraudScore     int              `json:"fraudScore"`
	RiskLevel      models.RiskLevel `json:"riskLevel"`
	CODBlocked     bool             `json:"codBlocked"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}

func ToReceiverFraudProfileResponse(r *models.ReceiverFraudProfile) ReceiverFraudProfileResponse {
	return ReceiverFraudProfileResponse{
		ID:             r.ID,
		Name:           r.Name,
		Phone:          r.Phone,
		Email:          r.Email,
		TotalOrders:    r.TotalOrders,
		TotalDelivered: r.TotalDelivered,
		FraudScore:     r.FraudScore,
		RiskLevel:      r.RiskLevel,
		CODBlocked:     r.CODBlocked,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}
