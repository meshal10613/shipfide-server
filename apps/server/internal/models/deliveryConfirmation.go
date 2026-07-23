package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeliveryConfirmation struct {
	ID            string    `gorm:"type:uuid;primaryKey"`
	ShipmentID    string    `gorm:"type:uuid;uniqueIndex;not null"`
	OtpHash       string    `gorm:"not null"` // bcrypt hash — never stored plaintext
	OtpExpiresAt  time.Time `gorm:"not null"` // 48 hours after rider assignment
	WrongAttempts int       `gorm:"default:0"`
	IsLocked      bool      `gorm:"default:false"` // locked after 3 wrong attempts
	IsUsed        bool      `gorm:"default:false"`
	UsedAt        *time.Time
	// Admin can regenerate once; reason must be logged separately in AuditLog.
	RegeneratedAt      *time.Time
	RegeneratedByAdmin *string `gorm:"type:uuid"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Shipment *Shipment `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE"`
}

func (d *DeliveryConfirmation) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.NewString()
	}
	return nil
}
