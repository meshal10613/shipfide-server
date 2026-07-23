package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type WalkInPaymentResponse struct {
	ID                 string                     `json:"id"`
	ShipmentID         string                     `json:"shipmentId"`
	Amount             decimal.Decimal            `json:"amount"`
	Method             models.PaymentMethod       `json:"method"`
	Status             models.WalkInPaymentStatus `json:"status"`
	CollectedByAdminID string                     `json:"collectedByAdminId"`
	CollectedAt        time.Time                  `json:"collectedAt"`
	CreatedAt          time.Time                  `json:"createdAt"`
}

func ToWalkInPaymentResponse(w *models.WalkInPayment) WalkInPaymentResponse {
	return WalkInPaymentResponse{
		ID:                 w.ID,
		ShipmentID:         w.ShipmentID,
		Amount:             w.Amount,
		Method:             w.Method,
		Status:             w.Status,
		CollectedByAdminID: w.CollectedByAdminID,
		CollectedAt:        w.CollectedAt,
		CreatedAt:          w.CreatedAt,
	}
}
