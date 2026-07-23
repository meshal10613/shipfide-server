package pricing

import (
	"testing"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

func TestCalculateDeliveryCharge(t *testing.T) {
	req := PricingRequest{
		Zone:        models.ZoneInsideDhaka,
		WeightGrams: 1500, // 1000g base + 500g extra = 60 + 15 = 75 BDT
		ParcelType:  models.ParcelPackage,
		IsFragile:   false,
	}

	res := CalculateDeliveryCharge(req)

	expected := decimal.NewFromInt(75)
	if !res.DeliveryCharge.Equal(expected) {
		t.Errorf("expected delivery charge %s, got %s", expected, res.DeliveryCharge)
	}
}

func TestCalculateRevenueSplit(t *testing.T) {
	deliveryCharge := decimal.NewFromInt(100)
	codAmount := decimal.NewFromInt(1000)

	res := CalculateRevenueSplit(models.SplitSameCity, deliveryCharge, codAmount)

	if !res.RiderShare.Equal(decimal.NewFromInt(50)) {
		t.Errorf("expected rider share 50, got %s", res.RiderShare)
	}
	if !res.SystemShare.Equal(decimal.NewFromInt(50)) {
		t.Errorf("expected system share 50, got %s", res.SystemShare)
	}
	if !res.MerchantNet.Equal(decimal.NewFromInt(900)) {
		t.Errorf("expected merchant net 900, got %s", res.MerchantNet)
	}
}
