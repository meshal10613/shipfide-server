package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin is a system staff member scoped to a hub.
type Admin struct {
	ID     string `gorm:"type:uuid;primaryKey"`
	UserID string `gorm:"type:uuid;not null;uniqueIndex"` // 1-to-1 with User

	Name  string `gorm:"not null"`
	Phone string `gorm:"type:varchar(20);not null"`

	// Geographic scope
	Division Division `gorm:"type:varchar(20);not null"`
	District District `gorm:"type:varchar(60);not null"`

	// A hub can have several staff (hub manager + operators), so HubID is a
	// plain (non-unique) index, not a 1-to-1 constraint.
	HubID *string `gorm:"type:uuid;index"`
	Hub   *Hub    `gorm:"foreignKey:HubID;references:ID"`

	CanApproveCashout bool `gorm:"default:true"`
	CanManageRiders   bool `gorm:"default:true"`
	CanViewFraudFlags bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Deleting or updating the parent User cascades to this Admin profile.
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// validateDivision checks that District belongs to Division, and auto-fills
// Division if it is blank.
func (a *Admin) validateDivision() error {
	if a.District == "" {
		return nil // super-admins may have no district
	}
	expected, ok := DivisionOf(a.District)
	if !ok {
		return fmt.Errorf("unknown district: %q", a.District)
	}
	if a.Division != "" && a.Division != expected {
		return fmt.Errorf(
			"district %q belongs to division %q, not %q",
			a.District, expected, a.Division,
		)
	}
	a.Division = expected
	return nil
}

func (a *Admin) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	return a.validateDivision()
}

func (a *Admin) BeforeSave(_ *gorm.DB) error {
	return a.validateDivision()
}
