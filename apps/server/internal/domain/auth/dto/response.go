package authDto

import "server/internal/models"

type AuthResponse struct {
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
	SessionToken string       `json:"sessionToken"`
	User         *models.User `json:"user"`
}
