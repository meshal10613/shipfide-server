package adminDto

import (
	"server/internal/models"
)

type CreateAdminRequest struct {
	Email             string          `json:"email" validate:"required,email"`
	Name              string          `json:"name" validate:"required,min=3"`
	Phone             string          `json:"phone" validate:"required"`
	District          models.District `json:"district" validate:"required"`
	HubID             *string         `json:"hubId" validate:"omitempty,uuid"`
	CanApproveCashout *bool           `json:"canApproveCashout" validate:"omitempty"`
	CanManageRiders   *bool           `json:"canManageRiders" validate:"omitempty"`
	CanViewFraudFlags *bool           `json:"canViewFraudFlags" validate:"omitempty"`
}

type UpdateAdminRequest struct {
	Name              *string          `json:"name" validate:"omitempty,min=3"`
	Phone             *string          `json:"phone" validate:"omitempty"`
	District          *models.District `json:"district" validate:"omitempty"`
	HubID             **string         `json:"hubId" validate:"omitempty"` // double pointer to support clearing HubID (nil)
	CanApproveCashout *bool            `json:"canApproveCashout" validate:"omitempty"`
	CanManageRiders   *bool            `json:"canManageRiders" validate:"omitempty"`
	CanViewFraudFlags *bool            `json:"canViewFraudFlags" validate:"omitempty"`
}
