package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type WithdrawalResponse struct {
	ID                    string                  `json:"id"`
	RiderID               *string                 `json:"riderId,omitempty"`
	MerchantID            *string                 `json:"merchantId,omitempty"`
	Amount                decimal.Decimal         `json:"amount"`
	Method                models.PaymentMethod    `json:"method"`
	AccountNumber         string                  `json:"accountNumber,omitempty"`
	BankName              string                  `json:"bankName,omitempty"`
	Status                models.WithdrawalStatus `json:"status"`
	RejectionReason       string                  `json:"rejectionReason,omitempty"`
	HubNotified           bool                    `json:"hubNotified"`
	HubIDVerified         bool                    `json:"hubIdVerified"`
	SendEmailConfirmation bool                    `json:"sendEmailConfirmation"`
	RequestedAt           time.Time               `json:"requestedAt"`
	ApprovedAt            *time.Time              `json:"approvedAt,omitempty"`
	PaidAt                *time.Time              `json:"paidAt,omitempty"`
}

func ToWithdrawalResponse(w *models.Withdrawal) WithdrawalResponse {
	return WithdrawalResponse{
		ID:                    w.ID,
		RiderID:               w.RiderID,
		MerchantID:            w.MerchantID,
		Amount:                w.Amount,
		Method:                w.Method,
		AccountNumber:         w.AccountNumber,
		BankName:              w.BankName,
		Status:                w.Status,
		RejectionReason:       w.RejectionReason,
		HubNotified:           w.HubNotified,
		HubIDVerified:         w.HubIDVerified,
		SendEmailConfirmation: w.SendEmailConfirmation,
		RequestedAt:           w.RequestedAt,
		ApprovedAt:            w.ApprovedAt,
		PaidAt:                w.PaidAt,
	}
}
