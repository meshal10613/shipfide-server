package session

import (
	"time"

	"gorm.io/gorm"
	"server/internal/models"
)

type Repository interface {
	FindActiveSessionsByUserID(userID string) ([]models.Session, error)
	FindByID(id string) (*models.Session, error)
	DeleteByID(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindActiveSessionsByUserID(userID string) ([]models.Session, error) {
	var sessions []models.Session
	err := r.db.Where("user_id = ? AND revoked_at IS NULL AND expires_at > ?", userID, time.Now()).
		Order("created_at DESC").
		Find(&sessions).
		Error
	return sessions, err
}

func (r *repository) FindByID(id string) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *repository) DeleteByID(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Session{}).Error
}
