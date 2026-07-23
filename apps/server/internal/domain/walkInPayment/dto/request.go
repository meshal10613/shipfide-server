package dto

import (
	"server/internal/models"

	"github.com/shopspring/decimal"
)

type CreateWalkInPaymentRequest struct {
	ShipmentID string               `json:"shipmentId" validate:"required"`
	Amount     decimal.Decimal      `json:"amount" validate:"required"`
	Method     models.PaymentMethod `json:"method" validate:"required"`
}
