package rider

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(rider *models.Rider) error
	GetByID(id string) (*models.Rider, error)
	GetByUserID(userID string) (*models.Rider, error)
	Update(rider *models.Rider) error
	List() ([]models.Rider, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(rider *models.Rider) error {
	return r.db.Create(rider).Error
}

func (r *repository) GetByID(id string) (*models.Rider, error) {
	var rider models.Rider
	err := r.db.Preload("User").Preload("AssignedHub").First(&rider, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rider, nil
}

func (r *repository) GetByUserID(userID string) (*models.Rider, error) {
	var rider models.Rider
	err := r.db.Preload("User").Preload("AssignedHub").First(&rider, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rider, nil
}

func (r *repository) Update(rider *models.Rider) error {
	return r.db.Save(rider).Error
}

func (r *repository) List() ([]models.Rider, error) {
	var riders []models.Rider
	err := r.db.Preload("User").Preload("AssignedHub").Order("created_at desc").Find(&riders).Error
	return riders, err
}
