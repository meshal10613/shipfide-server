package dto

import (
	"server/internal/models"

	"github.com/shopspring/decimal"
)

type CreateWithdrawalRequest struct {
	RiderID               *string              `json:"riderId,omitempty"`
	MerchantID            *string              `json:"merchantId,omitempty"`
	Amount                decimal.Decimal      `json:"amount" validate:"required"`
	Method                models.PaymentMethod `json:"method" validate:"required"`
	AccountNumber         string               `json:"accountNumber" validate:"omitempty"`
	BankName              string               `json:"bankName" validate:"omitempty"`
	SendEmailConfirmation bool                 `json:"sendEmailConfirmation"`
}

type UpdateWithdrawalStatusRequest struct {
	Status          models.WithdrawalStatus `json:"status" validate:"required"`
	RejectionReason string                  `json:"rejectionReason,omitempty"`
}
