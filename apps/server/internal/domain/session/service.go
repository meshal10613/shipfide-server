package session

import (
	"errors"

	"server/internal/domain/session/dto"
)

type Service interface {
	GetUserSessions(userID string) ([]*sessionDto.SessionResponse, error)
	DeleteSession(sessionID string, callerID string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserSessions(userID string) ([]*sessionDto.SessionResponse, error) {
	sessions, err := s.repo.FindActiveSessionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return sessionDto.MapToSessionResponseList(sessions), nil
}

func (s *service) DeleteSession(sessionID string, callerID string) error {
	sess, err := s.repo.FindByID(sessionID)
	if err != nil {
		return errors.New("session not found")
	}

	// Enforce: only own session can be deleted
	if sess.UserID != callerID {
		return errors.New("forbidden: you can only delete your own sessions")
	}

	return s.repo.DeleteByID(sessionID)
}
