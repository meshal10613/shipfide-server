package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Merchant struct {
	ID     string `gorm:"type:uuid;primaryKey"`
	UserID string `gorm:"type:uuid;not null;uniqueIndex"` // 1-to-1 with User

	Name        string `gorm:"not null"`
	FathersName string `gorm:"not null"`
	MothersName string `gorm:"not null"`
	DateOfBirth time.Time
	Gender      Gender     `gorm:"type:varchar(20);not null"`
	BloodGroup  BloodGroup `gorm:"type:varchar(20);not null"`

	// Primary business location — Division is auto-filled from District.
	Division Division `gorm:"type:varchar(20)"`
	District District `gorm:"type:varchar(60)"`

	// Pickup location (may differ from business address).
	PickupDivision Division `gorm:"type:varchar(20)"`
	PickupDistrict District `gorm:"type:varchar(60)"`

	// Full address text (free-form, for display).
	Address string `gorm:"type:text"`

	IsKycVerified bool   `gorm:"default:false"`
	Nid           string `gorm:"type:text"`

	CodEnabled                bool            `gorm:"default:true"`
	MaxCodAmount              decimal.Decimal `gorm:"type:decimal(10,2);not null;default:50000"`
	HasCompletedFirstDelivery bool            `gorm:"default:false;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Deleting or updating the parent User cascades to this Merchant profile.
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Shipments               []Shipment               `gorm:"foreignKey:MerchantID;references:ID"`
	MerchantDeliveryRatings []MerchantDeliveryRating `gorm:"foreignKey:MerchantID;references:ID"`
}

func (m *Merchant) validateDivisions() error {
	if m.District != "" {
		exp, ok := DivisionOf(m.District)
		if !ok {
			return fmt.Errorf("unknown district: %q", m.District)
		}
		if m.Division != "" && m.Division != exp {
			return fmt.Errorf("district %q belongs to division %q, not %q", m.District, exp, m.Division)
		}
		m.Division = exp
	}
	if m.PickupDistrict != "" {
		exp, ok := DivisionOf(m.PickupDistrict)
		if !ok {
			return fmt.Errorf("unknown pickup district: %q", m.PickupDistrict)
		}
		if m.PickupDivision != "" && m.PickupDivision != exp {
			return fmt.Errorf("pickup district %q belongs to division %q, not %q", m.PickupDistrict, exp, m.PickupDivision)
		}
		m.PickupDivision = exp
	}
	return nil
}

func (m *Merchant) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	return m.validateDivisions()
}

func (m *Merchant) BeforeSave(_ *gorm.DB) error {
	return m.validateDivisions()
}
