package userDto

import (
	"time"
	"server/internal/models"
)

type UserResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Phone     *string           `json:"phone,omitempty"`
	Image     *string           `json:"image,omitempty"`
	Role      models.Role       `json:"role"`
	Status    models.UserStatus `json:"status"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func MapToUserResponse(u *models.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Image:     u.Image,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func MapToUserResponseList(users []models.User) []*UserResponse {
	res := make([]*UserResponse, len(users))
	for i, u := range users {
		res[i] = MapToUserResponse(&u)
	}
	return res
}
