package sessionDto

import (
	"time"
	"server/internal/models"
)

type SessionResponse struct {
	ID           string     `json:"id"`
	UserAgent    *string    `json:"userAgent,omitempty"`
	IPAddress    *string    `json:"ipAddress,omitempty"`
	DeviceName   *string    `json:"deviceName,omitempty"`
	ExpiresAt    time.Time  `json:"expiresAt"`
	RevokedAt    *time.Time `json:"revokedAt,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
}

func MapToSessionResponse(s *models.Session) *SessionResponse {
	return &SessionResponse{
		ID:           s.ID,
		UserAgent:    s.UserAgent,
		IPAddress:    s.IPAddress,
		DeviceName:   s.DeviceName,
		ExpiresAt:    s.ExpiresAt,
		RevokedAt:    s.RevokedAt,
		CreatedAt:    s.CreatedAt,
	}
}

func MapToSessionResponseList(sessions []models.Session) []*SessionResponse {
	res := make([]*SessionResponse, len(sessions))
	for i, s := range sessions {
		res[i] = MapToSessionResponse(&s)
	}
	return res
}
