package auth

import (
	"time"

	"gorm.io/gorm"
	"server/internal/models"
)

type Repository interface {
	CreateUser(user *models.User, account *models.Account) error
	UpdateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByIDWithSessions(id string) (*models.User, error)
	GetAccountByUserID(userID string) (*models.Account, error)
	GetAccountByRefreshToken(refreshToken string) (*models.Account, error)
	UpdateAccount(account *models.Account) error
	CreateSession(session *models.Session) error
	UpdateSession(session *models.Session) error
	GetSessionByToken(sessionToken string) (*models.Session, error)
	DeleteSessionByToken(sessionToken string) error
	GetActiveSessionsCount(userID string) (int64, error)
	DeleteAllSessionsByUserID(userID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *models.User, account *models.Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		account.UserID = user.ID
		if err := tx.Create(account).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Account").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Account").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByIDWithSessions(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Account").Preload("Sessions").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetAccountByUserID(userID string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *repository) GetAccountByRefreshToken(refreshToken string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *repository) UpdateAccount(account *models.Account) error {
	return r.db.Save(account).Error
}

func (r *repository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *repository) UpdateSession(session *models.Session) error {
	return r.db.Save(session).Error
}

func (r *repository) GetSessionByToken(sessionToken string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("session_token = ?", sessionToken).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *repository) DeleteSessionByToken(sessionToken string) error {
	return r.db.Where("session_token = ?", sessionToken).Delete(&models.Session{}).Error
}

func (r *repository) GetActiveSessionsCount(userID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Session{}).
		Where("user_id = ? AND revoked_at IS NULL AND expires_at > ?", userID, time.Now()).
		Count(&count).
		Error
	return count, err
}

func (r *repository) DeleteAllSessionsByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Session{}).Error
}
