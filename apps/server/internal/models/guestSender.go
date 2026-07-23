package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GuestSender struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"not null"`

	// Phone is the real identity key for a walk-in customer — it's always
	// collected, unlike email. A previous version required a unique Email
	// instead, which broke on the second guest with no email (two blank
	// strings collide on a NOT NULL UNIQUE column).
	Phone string  `gorm:"type:varchar(20);not null;uniqueIndex"`
	Email *string `gorm:"type:text;uniqueIndex"` // optional — multiple NULLs are allowed by a unique index

	// Location — Division auto-filled from District.
	Division Division `gorm:"type:varchar(20)"`
	District District `gorm:"type:varchar(60)"`

	// Full address text (for display / high-value parcels).
	Address   *string `gorm:"type:text"`
	NidNumber *string `gorm:"type:text"` // optional — required for parcels >50,000 BDT

	IsPhoneFlagged bool   `gorm:"default:false"`
	FlagReason     string `gorm:"type:varchar(255)"`

	AdminID string `gorm:"type:uuid;not null;index"` // which admin registered this guest
	Admin   *Admin `gorm:"foreignKey:AdminID;references:ID"`

	CreatedAt time.Time

	Shipments []Shipment `gorm:"foreignKey:GuestSenderID"`
}

func (g *GuestSender) validateDivision() error {
	if g.District == "" {
		return nil
	}
	exp, ok := DivisionOf(g.District)
	if !ok {
		return fmt.Errorf("unknown district: %q", g.District)
	}
	if g.Division != "" && g.Division != exp {
		return fmt.Errorf("district %q belongs to division %q, not %q", g.District, exp, g.Division)
	}
	g.Division = exp
	return nil
}

func (g *GuestSender) BeforeCreate(tx *gorm.DB) error {
	if g.ID == "" {
		g.ID = uuid.NewString()
	}
	return g.validateDivision()
}

func (g *GuestSender) BeforeSave(_ *gorm.DB) error {
	return g.validateDivision()
}
