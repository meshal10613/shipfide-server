package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CodDeliveryConfirmation struct {
	ID         string `gorm:"type:uuid;primaryKey"`
	ShipmentID string `gorm:"type:uuid;uniqueIndex;not null"`
	// ExpectedAmount must equal codAmount on the Shipment.
	ExpectedAmount decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	// CollectedAmount is entered by the rider at delivery — must match
	// ExpectedAmount exactly. This is validated in BeforeSave below; the
	// comment previously promised this but nothing enforced it.
	CollectedAmount *decimal.Decimal `gorm:"type:decimal(10,2)"`
	ConfirmedAt     *time.Time
	// DepositedAt is set when rider submits cash at hub (within 24h deadline).
	DepositedAt               *time.Time
	DepositConfirmedByAdminID *string `gorm:"type:uuid;index"`
	DepositConfirmedAt        *time.Time
	CreatedAt                 time.Time
	UpdatedAt                 time.Time

	Shipment                *Shipment `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE"`
	DepositConfirmedByAdmin *Admin    `gorm:"foreignKey:DepositConfirmedByAdminID"`
}

func (c *CodDeliveryConfirmation) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.NewString()
	}
	return nil
}

// BeforeSave enforces that, once a rider records CollectedAmount, it matches
// ExpectedAmount exactly — a COD collection short (or over) by even one taka
// must be resolved as an explicit exception, not silently persisted.
func (c *CodDeliveryConfirmation) BeforeSave(_ *gorm.DB) error {
	if c.CollectedAmount != nil && !c.CollectedAmount.Equal(c.ExpectedAmount) {
		return fmt.Errorf(
			"collected amount %s does not match expected amount %s",
			c.CollectedAmount.String(), c.ExpectedAmount.String(),
		)
	}
	return nil
}
