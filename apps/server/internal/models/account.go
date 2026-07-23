package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Account holds authentication-related credentials and tokens for a User.
// It has a strict one-to-one relationship with User (one Account per User).
type Account struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`

	AccessToken           *string    `gorm:"type:text" json:"accessToken,omitempty"`
	RefreshToken          *string    `gorm:"type:text" json:"refreshToken,omitempty"`
	AccessTokenAt         *time.Time `json:"accessTokenAt,omitempty"` // when the current access token was issued
	RefreshTokenExpiresAt *time.Time `json:"refreshTokenExpiresAt,omitempty"`

	Password string `gorm:"not null" json:"password"`

	Otp          *string    `gorm:"type:varchar(10)" json:"otp,omitempty"`
	OtpExpiresAt *time.Time `json:"otpExpiresAt,omitempty"`

	UserID string `gorm:"type:uuid;not null;uniqueIndex" json:"userId"` // uniqueIndex enforces one-to-one with User
	User   *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"user,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	return nil
}

// HashPassword hashes a plaintext password and sets it on the Account.
func (a *Account) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}

// CheckPassword verifies the plaintext password against the stored hashed password.
func (a *Account) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}
