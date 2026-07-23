package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalkInPayment struct {
	ID                 string              `gorm:"type:uuid;primaryKey"`
	ShipmentID         string              `gorm:"type:uuid;uniqueIndex;not null"`
	Amount             decimal.Decimal     `gorm:"type:decimal(10,2);not null"`
	Method             PaymentMethod       `gorm:"type:varchar(20);not null;default:'CASH'"`
	Status             WalkInPaymentStatus `gorm:"type:varchar(20);not null;default:'COLLECTED'"`
	CollectedByAdminID string              `gorm:"type:uuid;not null;index"`
	CollectedAt        time.Time           `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Shipment         *Shipment `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE"`
	CollectedByAdmin *Admin    `gorm:"foreignKey:CollectedByAdminID"`
}

func (w *WalkInPayment) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	if w.CollectedAt.IsZero() {
		w.CollectedAt = time.Now()
	}
	return nil
}
