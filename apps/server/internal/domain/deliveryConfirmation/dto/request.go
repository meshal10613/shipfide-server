package dto

import "github.com/shopspring/decimal"

type VerifyOtpRequest struct {
	ShipmentID string `json:"shipmentId" validate:"required"`
	Otp        string `json:"otp" validate:"required,len=6"`
}

type RegenerateOtpRequest struct {
	ShipmentID string `json:"shipmentId" validate:"required"`
	Reason     string `json:"reason" validate:"required"`
}

type ConfirmCodRequest struct {
	ShipmentID      string          `json:"shipmentId" validate:"required"`
	CollectedAmount decimal.Decimal `json:"collectedAmount" validate:"required"`
}

type DepositCodRequest struct {
	CodConfirmationIDs []string `json:"codConfirmationIds" validate:"required"`
}
