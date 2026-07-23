package withdrawal

import (
	"testing"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

func TestWithdrawalValidation(t *testing.T) {
	riderID := "rider-123"
	merchantID := "merchant-456"

	// Invalid: both rider and merchant set
	wInvalid := &models.Withdrawal{
		RiderID:    &riderID,
		MerchantID: &merchantID,
		Amount:     decimal.NewFromInt(500),
		Method:     models.PaymentBkash,
	}

	if err := wInvalid.ValidateParty(); err == nil {
		t.Errorf("expected error when both RiderID and MerchantID are set")
	}

	// Invalid: amount < 100
	wLow := &models.Withdrawal{
		RiderID: &riderID,
		Amount:  decimal.NewFromInt(50),
		Method:  models.PaymentBkash,
	}
	if err := wLow.ValidateMinimum(); err == nil {
		t.Errorf("expected error when withdrawal amount is less than 100")
	}

	// Valid
	wValid := &models.Withdrawal{
		RiderID: &riderID,
		Amount:  decimal.NewFromInt(200),
		Method:  models.PaymentBkash,
	}
	if err := wValid.ValidateParty(); err != nil {
		t.Errorf("unexpected party validation error: %v", err)
	}
	if err := wValid.ValidateMinimum(); err != nil {
		t.Errorf("unexpected minimum validation error: %v", err)
	}
}
