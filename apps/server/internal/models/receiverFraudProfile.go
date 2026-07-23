package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ReceiverFraudProfile tracks delivery success rate per receiver phone number.
// A new profile starts at score 100 (trusted) and degrades as failed deliveries accumulate.
type ReceiverFraudProfile struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"type:varchar(120)"`

	// Phone is the real lookup key: every shipment has a required ReceiverPhone
	// but only an optional ReceiverEmail, so fraud checks at shipment-creation
	// time key off phone. Email is kept as an optional secondary identifier.
	Phone string  `gorm:"type:varchar(20);not null;uniqueIndex"`
	Email *string `gorm:"uniqueIndex"`

	TotalOrders    int       `gorm:"default:0"`
	TotalDelivered int       `gorm:"default:0"`
	FraudScore     int       `gorm:"default:100"` // 0–100
	RiskLevel      RiskLevel `gorm:"type:varchar(20);default:'NEW_CUSTOMER'"`
	CODBlocked     bool      `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// RefreshScore recomputes FraudScore, RiskLevel, and CODBlocked.
// Call this after updating TotalOrders / TotalDelivered.
func (r *ReceiverFraudProfile) RefreshScore() {
	if r.TotalOrders == 0 {
		r.FraudScore = 100
	} else {
		r.FraudScore = int(float64(r.TotalDelivered) / float64(r.TotalOrders) * 100)
	}
	switch {
	case r.TotalOrders == 0:
		r.RiskLevel = RiskLevelNewCustomer
	case r.FraudScore >= 81:
		r.RiskLevel = RiskLevelTrusted
	case r.FraudScore >= 61:
		r.RiskLevel = RiskLevelLow
	case r.FraudScore >= 41:
		r.RiskLevel = RiskLevelMedium
	case r.FraudScore >= 21:
		r.RiskLevel = RiskLevelHigh
	default:
		r.RiskLevel = RiskLevelBlacklisted
	}
	r.CODBlocked = r.FraudScore <= 20
}

func (r *ReceiverFraudProfile) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	r.RefreshScore()
	return nil
}
