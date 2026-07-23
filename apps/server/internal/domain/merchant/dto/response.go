package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type MerchantResponse struct {
	ID                        string            `json:"id"`
	UserID                    string            `json:"userId"`
	Name                      string            `json:"name"`
	FathersName               string            `json:"fathersName"`
	MothersName               string            `json:"mothersName"`
	DateOfBirth               time.Time         `json:"dateOfBirth"`
	Gender                    models.Gender     `json:"gender"`
	BloodGroup                models.BloodGroup `json:"bloodGroup"`
	Division                  models.Division   `json:"division"`
	District                  models.District   `json:"district"`
	PickupDivision            models.Division   `json:"pickupDivision,omitempty"`
	PickupDistrict            models.District   `json:"pickupDistrict,omitempty"`
	Address                   string            `json:"address,omitempty"`
	IsKycVerified             bool              `json:"isKycVerified"`
	Nid                       string            `json:"nid,omitempty"`
	CodEnabled                bool              `json:"codEnabled"`
	MaxCodAmount              decimal.Decimal   `json:"maxCodAmount"`
	HasCompletedFirstDelivery bool              `json:"hasCompletedFirstDelivery"`
	CreatedAt                 time.Time         `json:"createdAt"`
	UpdatedAt                 time.Time         `json:"updatedAt"`
}

func ToMerchantResponse(m *models.Merchant) MerchantResponse {
	return MerchantResponse{
		ID:                        m.ID,
		UserID:                    m.UserID,
		Name:                      m.Name,
		FathersName:               m.FathersName,
		MothersName:               m.MothersName,
		DateOfBirth:               m.DateOfBirth,
		Gender:                    m.Gender,
		BloodGroup:                m.BloodGroup,
		Division:                  m.Division,
		District:                  m.District,
		PickupDivision:            m.PickupDivision,
		PickupDistrict:            m.PickupDistrict,
		Address:                   m.Address,
		IsKycVerified:             m.IsKycVerified,
		Nid:                       m.Nid,
		CodEnabled:                m.CodEnabled,
		MaxCodAmount:              m.MaxCodAmount,
		HasCompletedFirstDelivery: m.HasCompletedFirstDelivery,
		CreatedAt:                 m.CreatedAt,
		UpdatedAt:                 m.UpdatedAt,
	}
}
