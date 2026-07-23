package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type CreateMerchantRequest struct {
	UserID         string            `json:"userId" validate:"omitempty"` // populated automatically if self
	Name           string            `json:"name" validate:"required"`
	FathersName    string            `json:"fathersName" validate:"required"`
	MothersName    string            `json:"mothersName" validate:"required"`
	DateOfBirth    time.Time         `json:"dateOfBirth" validate:"required"`
	Gender         models.Gender     `json:"gender" validate:"required"`
	BloodGroup     models.BloodGroup `json:"bloodGroup" validate:"required"`
	District       models.District   `json:"district" validate:"required"`
	Division       models.Division   `json:"division" validate:"omitempty"`
	PickupDistrict models.District   `json:"pickupDistrict" validate:"omitempty"`
	PickupDivision models.Division   `json:"pickupDivision" validate:"omitempty"`
	Address        string            `json:"address" validate:"omitempty"`
	Nid            string            `json:"nid" validate:"omitempty"`
}

type UpdateMerchantRequest struct {
	Name           string          `json:"name" validate:"omitempty"`
	FathersName    string          `json:"fathersName" validate:"omitempty"`
	MothersName    string          `json:"mothersName" validate:"omitempty"`
	District       models.District `json:"district" validate:"omitempty"`
	Division       models.Division `json:"division" validate:"omitempty"`
	PickupDistrict models.District `json:"pickupDistrict" validate:"omitempty"`
	PickupDivision models.Division `json:"pickupDivision" validate:"omitempty"`
	Address        string          `json:"address" validate:"omitempty"`
	Nid            string          `json:"nid" validate:"omitempty"`
}

type UpdateKycRequest struct {
	IsKycVerified bool             `json:"isKycVerified"`
	CodEnabled    *bool            `json:"codEnabled,omitempty"`
	MaxCodAmount  *decimal.Decimal `json:"maxCodAmount,omitempty"`
}
