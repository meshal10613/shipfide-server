package receiverFraud

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetByPhone(phone string) (*models.ReceiverFraudProfile, error)
	GetByID(id string) (*models.ReceiverFraudProfile, error)
	Update(profile *models.ReceiverFraudProfile) error
	List() ([]models.ReceiverFraudProfile, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByPhone(phone string) (*models.ReceiverFraudProfile, error) {
	var profile models.ReceiverFraudProfile
	err := r.db.First(&profile, "phone = ?", phone).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *repository) GetByID(id string) (*models.ReceiverFraudProfile, error) {
	var profile models.ReceiverFraudProfile
	err := r.db.First(&profile, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *repository) Update(profile *models.ReceiverFraudProfile) error {
	return r.db.Save(profile).Error
}

func (r *repository) List() ([]models.ReceiverFraudProfile, error) {
	var profiles []models.ReceiverFraudProfile
	err := r.db.Order("updated_at desc").Find(&profiles).Error
	return profiles, err
}
