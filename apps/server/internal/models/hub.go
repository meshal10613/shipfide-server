package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hub represents a physical operations hub.
// A super admin creates hubs; each hub can have several staff (Admin) assigned to it.
type Hub struct {
	ID string `gorm:"type:uuid;primaryKey"`
	// Name uniqueness is scoped to non-deleted hubs so a closed hub's name
	// can be reused by a hub opened later at the same or a different location.
	Name string `gorm:"type:varchar(120);not null;uniqueIndex:idx_hubs_name,where:deleted_at IS NULL"`

	Division   Division `gorm:"type:varchar(20);not null"`
	District   District `gorm:"type:varchar(60);not null"`
	PostalCode string   `gorm:"type:varchar(10);not null"`
	Address    string   `gorm:"type:text;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Staff assigned to this hub (hub manager + operators). Deleting a hub does
	// NOT cascade-delete staff — a hub closure should not delete Admin/User
	// accounts; reassign staff to another hub first.
	Staff []Admin `gorm:"foreignKey:HubID;references:ID"`
}

// validateDivision checks that District belongs to Division, and auto-fills
// Division if it is blank.
func (h *Hub) validateDivision() error {
	if h.District == "" {
		return fmt.Errorf("district is required")
	}
	expected, ok := DivisionOf(h.District)
	if !ok {
		return fmt.Errorf("unknown district: %q", h.District)
	}
	if h.Division != "" && h.Division != expected {
		return fmt.Errorf(
			"district %q belongs to division %q, not %q",
			h.District, expected, h.Division,
		)
	}
	h.Division = expected
	return nil
}

func (h *Hub) BeforeCreate(tx *gorm.DB) error {
	if h.ID == "" {
		h.ID = uuid.NewString()
	}
	return h.validateDivision()
}

func (h *Hub) BeforeSave(_ *gorm.DB) error {
	return h.validateDivision()
}
