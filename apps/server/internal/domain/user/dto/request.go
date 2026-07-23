package userDto

import "server/internal/models"

type UpdateUserRequest struct {
	Name   *string            `json:"name" validate:"omitempty,min=3"`
	Phone  *string            `json:"phone" validate:"omitempty"`
	Image  *string            `json:"image" validate:"omitempty"`
	Role   *models.Role       `json:"role" validate:"omitempty"`
	Status *models.UserStatus `json:"status" validate:"omitempty"`
}
