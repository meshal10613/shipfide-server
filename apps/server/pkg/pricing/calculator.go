package pricing

import (
	"math"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type PricingRequest struct {
	Zone            models.ZoneType
	WeightGrams     int
	ParcelType      models.ParcelType
	IsFragile       bool
	PickupSurcharge *decimal.Decimal
}

type PricingResult struct {
	DeliveryCharge  decimal.Decimal
	PickupSurcharge decimal.Decimal
	TotalCharge     decimal.Decimal
}

// CalculateDeliveryCharge calculates the total shipping fees based on zone, weight, parcel type, and surcharges.
func CalculateDeliveryCharge(req PricingRequest) PricingResult {
	var baseCharge decimal.Decimal
	var addPer500g decimal.Decimal

	switch req.Zone {
	case models.ZoneInsideDhaka:
		baseCharge = decimal.NewFromInt(60)
		addPer500g = decimal.NewFromInt(15)
	case models.ZoneDhakaSuburb:
		baseCharge = decimal.NewFromInt(80)
		addPer500g = decimal.NewFromInt(15)
	case models.ZoneSameCityNonDhaka:
		baseCharge = decimal.NewFromInt(70)
		addPer500g = decimal.NewFromInt(15)
	case models.ZoneOutsideDhakaDivisional:
		baseCharge = decimal.NewFromInt(100)
		addPer500g = decimal.NewFromInt(25)
	case models.ZoneOutsideDhakaDistrict:
		baseCharge = decimal.NewFromInt(120)
		addPer500g = decimal.NewFromInt(25)
	case models.ZoneOutsideDhakaUpazila:
		baseCharge = decimal.NewFromInt(150)
		addPer500g = decimal.NewFromInt(25)
	default:
		// Default fallback
		baseCharge = decimal.NewFromInt(100)
		addPer500g = decimal.NewFromInt(20)
	}

	deliveryCharge := baseCharge

	// Additional weight charge for weight > 1000g
	if req.WeightGrams > 1000 {
		excessWeight := req.WeightGrams - 1000
		steps := int(math.Ceil(float64(excessWeight) / 500.0))
		weightAdd := addPer500g.Mul(decimal.NewFromInt(int64(steps)))
		deliveryCharge = deliveryCharge.Add(weightAdd)
	}

	// Parcel type / Fragile surcharge
	if req.IsFragile || req.ParcelType == models.ParcelFragile || req.ParcelType == models.ParcelHeavy {
		deliveryCharge = deliveryCharge.Add(decimal.NewFromInt(20))
	}

	surcharge := decimal.Zero
	if req.PickupSurcharge != nil && req.PickupSurcharge.IsPositive() {
		surcharge = *req.PickupSurcharge
	}

	total := deliveryCharge.Add(surcharge)

	return PricingResult{
		DeliveryCharge:  deliveryCharge,
		PickupSurcharge: surcharge,
		TotalCharge:     total,
	}
}

type RevenueSplitResult struct {
	SplitZone     models.SplitZone
	RiderSharePct decimal.Decimal
	RiderShare    decimal.Decimal
	SystemShare   decimal.Decimal
	MerchantNet   decimal.Decimal
}

// CalculateRevenueSplit computes rider share, system share, and merchant net payout upon delivery.
func CalculateRevenueSplit(
	splitZone models.SplitZone,
	deliveryCharge decimal.Decimal,
	codAmount decimal.Decimal,
) RevenueSplitResult {
	var riderPct decimal.Decimal

	switch splitZone {
	case models.SplitSameCity, models.SplitSameDistrict:
		riderPct = decimal.NewFromFloat(50.0) // 50%
	case models.SplitOutsideDistrict:
		riderPct = decimal.NewFromFloat(40.0) // 40%
	default:
		riderPct = decimal.NewFromFloat(50.0)
	}

	riderShare := deliveryCharge.Mul(riderPct).Div(decimal.NewFromInt(100)).Round(2)
	systemShare := deliveryCharge.Sub(riderShare)

	// Merchant net payout = codAmount - deliveryCharge (if codAmount > 0)
	merchantNet := codAmount.Sub(deliveryCharge)
	if merchantNet.IsNegative() {
		merchantNet = decimal.Zero
	}

	return RevenueSplitResult{
		SplitZone:     splitZone,
		RiderSharePct: riderPct,
		RiderShare:    riderShare,
		SystemShare:   systemShare,
		MerchantNet:   merchantNet,
	}
}
