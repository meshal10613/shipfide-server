package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"type:uuid;primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null;uniqueIndex" json:"email"`

	Phone *string `gorm:"type:varchar(20);uniqueIndex" json:"phone,omitempty"`
	Image *string `gorm:"type:text" json:"image,omitempty"`

	Role   Role       `gorm:"type:varchar(20);not null;default:'MERCHANT'" json:"role"`
	Status UserStatus `gorm:"type:varchar(20);not null;default:'PENDING'" json:"status"`

	NeedsPasswordChange bool `gorm:"default:false" json:"needsPasswordChange"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Account  *Account  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"`
	Sessions []Session `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"sessions,omitempty"`

	Admin    *Admin    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
