package dto

import "server/internal/models"

type CreateGuestSenderRequest struct {
	Name      string          `json:"name" validate:"required"`
	Phone     string          `json:"phone" validate:"required"`
	Email     *string         `json:"email" validate:"omitempty,email"`
	District  models.District `json:"district" validate:"omitempty"`
	Division  models.Division `json:"division" validate:"omitempty"`
	Address   *string         `json:"address" validate:"omitempty"`
	NidNumber *string         `json:"nidNumber" validate:"omitempty"`
}

type FlagGuestSenderRequest struct {
	IsPhoneFlagged bool   `json:"isPhoneFlagged"`
	FlagReason     string `json:"flagReason" validate:"omitempty"`
}
