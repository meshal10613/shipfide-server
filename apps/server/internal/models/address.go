package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Address is a reusable address record stored in its own table.
// Division is auto-filled from District on save; supplying a mismatched
// Division returns an error.
type Address struct {
	ID string `gorm:"type:uuid;primaryKey"`

	Zone     ZoneType `gorm:"type:varchar(40)"`
	Division Division `gorm:"type:varchar(20);not null"`
	District District `gorm:"type:varchar(60);not null"`

	FullAddress string `gorm:"type:text;not null"`
	AreaDetail  string `gorm:"type:varchar(255)"` // road / house / flat
	PostalCode  string `gorm:"type:varchar(10)"`
	CreatedAt   time.Time
}

// ValidateDivision checks that District belongs to the declared Division,
// and auto-fills Division if it is blank.
func (a *Address) ValidateDivision() error {
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

func (a *Address) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	return a.ValidateDivision()
}

func (a *Address) BeforeSave(_ *gorm.DB) error {
	return a.ValidateDivision()
}
