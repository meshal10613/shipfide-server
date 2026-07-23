package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Withdrawal represents a cashout request from a Rider or Merchant.
// Exactly one of RiderID / MerchantID must be set — enforced both at the
// application layer (ValidateParty, called from BeforeCreate) and at the
// database layer (CHECK constraint below), so a bug in a future code path
// that bypasses the Go hook still can't write an invalid row.
type Withdrawal struct {
	ID         string  `gorm:"type:uuid;primaryKey"`
	RiderID    *string `gorm:"type:uuid;index;check:chk_withdrawal_one_party,(rider_id IS NOT NULL AND merchant_id IS NULL) OR (rider_id IS NULL AND merchant_id IS NOT NULL)"`
	MerchantID *string `gorm:"type:uuid;index"`

	Amount        decimal.Decimal `gorm:"type:decimal(12,2);not null;check:chk_withdrawal_min_amount,amount >= 100"`
	Method        PaymentMethod   `gorm:"type:varchar(20);not null"`
	AccountNumber string          `gorm:"type:varchar(30)"`
	BankName      string          `gorm:"type:varchar(100)"`

	Status          WithdrawalStatus `gorm:"type:varchar(20);default:'REQUESTED'"`
	RejectionReason string           `gorm:"type:varchar(255)"`

	// Hub pickup fields
	HubNotified   bool `gorm:"default:false"`
	HubIDVerified bool `gorm:"default:false"`

	// Optional email confirmation
	SendEmailConfirmation bool `gorm:"default:false"`
	EmailSentAt           *time.Time

	RequestedAt time.Time `gorm:"autoCreateTime"`
	ApprovedAt  *time.Time
	PaidAt      *time.Time

	Rider    *Rider    `gorm:"foreignKey:RiderID"`
	Merchant *Merchant `gorm:"foreignKey:MerchantID"`
}

// ValidateMinimum returns an error if the withdrawal amount is below ৳100.
func (w *Withdrawal) ValidateMinimum() error {
	if w.Amount.LessThan(decimal.NewFromInt(100)) {
		return errors.New("minimum withdrawal amount is ৳100")
	}
	return nil
}

// ValidateParty returns an error unless exactly one of RiderID / MerchantID is set.
func (w *Withdrawal) ValidateParty() error {
	if (w.RiderID == nil) == (w.MerchantID == nil) {
		return errors.New("withdrawal must have exactly one of RiderID or MerchantID set")
	}
	return nil
}

func (w *Withdrawal) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	if err := w.ValidateParty(); err != nil {
		return err
	}
	return w.ValidateMinimum()
}
