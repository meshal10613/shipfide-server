package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type DeliveryConfirmationResponse struct {
	ID            string     `json:"id"`
	ShipmentID    string     `json:"shipmentId"`
	OtpExpiresAt  time.Time  `json:"otpExpiresAt"`
	WrongAttempts int        `json:"wrongAttempts"`
	IsLocked      bool       `json:"isLocked"`
	IsUsed        bool       `json:"isUsed"`
	UsedAt        *time.Time `json:"usedAt,omitempty"`
}

type CodConfirmationResponse struct {
	ID                 string           `json:"id"`
	ShipmentID         string           `json:"shipmentId"`
	ExpectedAmount     decimal.Decimal  `json:"expectedAmount"`
	CollectedAmount    *decimal.Decimal `json:"collectedAmount,omitempty"`
	ConfirmedAt        *time.Time       `json:"confirmedAt,omitempty"`
	DepositedAt        *time.Time       `json:"depositedAt,omitempty"`
	DepositConfirmedAt *time.Time       `json:"depositConfirmedAt,omitempty"`
}

func ToDeliveryConfirmationResponse(d *models.DeliveryConfirmation) DeliveryConfirmationResponse {
	return DeliveryConfirmationResponse{
		ID:            d.ID,
		ShipmentID:    d.ShipmentID,
		OtpExpiresAt:  d.OtpExpiresAt,
		WrongAttempts: d.WrongAttempts,
		IsLocked:      d.IsLocked,
		IsUsed:        d.IsUsed,
		UsedAt:        d.UsedAt,
	}
}

func ToCodConfirmationResponse(c *models.CodDeliveryConfirmation) CodConfirmationResponse {
	return CodConfirmationResponse{
		ID:                 c.ID,
		ShipmentID:         c.ShipmentID,
		ExpectedAmount:     c.ExpectedAmount,
		CollectedAmount:    c.CollectedAmount,
		ConfirmedAt:        c.ConfirmedAt,
		DepositedAt:        c.DepositedAt,
		DepositConfirmedAt: c.DepositConfirmedAt,
	}
}
