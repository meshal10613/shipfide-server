package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID           string    `gorm:"type:uuid;primaryKey" json:"id"`

	SessionToken string    `gorm:"not null" json:"sessionToken"`
	UserAgent    *string   `gorm:"type:text" json:"userAgent,omitempty"`
	IPAddress    *string   `gorm:"type:varchar(45)" json:"ipAddress,omitempty"` // supports IPv6
	DeviceName   *string   `gorm:"type:text" json:"deviceName,omitempty"`

	ExpiresAt    time.Time `gorm:"not null" json:"expiresAt"`
	RevokedAt    *time.Time `json:"revokedAt,omitempty"`

	UserID       string    `gorm:"type:uuid;not null;index" json:"userId"`
	User         *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"user,omitempty"`

	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.NewString()
	}
	return nil
}
