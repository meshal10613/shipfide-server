package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeliveryRating — receiver rates the rider after delivery (Phase 2).
// Defined here because Rider and Shipment already reference this type.
type DeliveryRating struct {
	ID              string  `gorm:"type:uuid;primaryKey"`
	ShipmentID      string  `gorm:"type:uuid;uniqueIndex;not null"`
	RiderID         string  `gorm:"type:uuid;not null;index"`
	ReceiverPhone   string  `gorm:"not null"`
	IsAnonymous     bool    `gorm:"default:false"`
	Stars           int     `gorm:"not null;check:chk_delivery_rating_stars,stars BETWEEN 1 AND 5"`
	Comment         *string `gorm:"type:text"` // max 500 chars
	IsVisible       bool    `gorm:"default:true"`
	IsFlagged       bool    `gorm:"default:false"`
	FlaggedReason   *string `gorm:"type:text"`
	ModeratedBy     *string `gorm:"type:uuid"`
	ModeratedAt     *time.Time
	RatingWindowEnd time.Time `gorm:"not null"` // 7 days from DELIVERED
	SubmittedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Shipment *Shipment `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE"`
	Rider    *Rider    `gorm:"foreignKey:RiderID"`
}

func (d *DeliveryRating) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.NewString()
	}
	return d.validateStars()
}

func (d *DeliveryRating) validateStars() error {
	if d.Stars < 1 || d.Stars > 5 {
		return fmt.Errorf("stars must be between 1 and 5, got %d", d.Stars)
	}
	return nil
}

// MerchantDeliveryRating — merchant rates the platform after delivery (Phase 2).
// Defined here because Merchant already references this type.
//
// Eligibility rules (enforced by the service layer):
//  1. One rating per shipment — guaranteed by the uniqueIndex on ShipmentID.
//  2. A MerchantDeliveryRating may only be created when Merchant.HasCompletedFirstDelivery == true.
//     The service must check this flag (set automatically when the first shipment reaches DELIVERED)
//     before allowing a rating submission.
type MerchantDeliveryRating struct {
	ID              string             `gorm:"type:uuid;primaryKey"`
	ShipmentID      string             `gorm:"type:uuid;uniqueIndex;not null"`
	MerchantID      string             `gorm:"type:uuid;not null;index"`
	Stars           int                `gorm:"not null;check:chk_merchant_rating_stars,stars BETWEEN 1 AND 5"`
	Comment         *string            `gorm:"type:text"` // max 1000 chars, admin-only
	Tag             *MerchantRatingTag `gorm:"type:varchar(30)"`
	RatingWindowEnd time.Time          `gorm:"not null"` // 14 days from DELIVERED/RETURNED_TO_MERCHANT
	SubmittedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Shipment *Shipment `gorm:"foreignKey:ShipmentID;constraint:OnDelete:CASCADE"`
	Merchant *Merchant `gorm:"foreignKey:MerchantID"`
}

func (m *MerchantDeliveryRating) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	if m.Stars < 1 || m.Stars > 5 {
		return fmt.Errorf("stars must be between 1 and 5, got %d", m.Stars)
	}
	return nil
}
