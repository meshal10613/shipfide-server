package merchant

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(merchant *models.Merchant) error
	GetByID(id string) (*models.Merchant, error)
	GetByUserID(userID string) (*models.Merchant, error)
	Update(merchant *models.Merchant) error
	List() ([]models.Merchant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(merchant *models.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *repository) GetByID(id string) (*models.Merchant, error) {
	var m models.Merchant
	err := r.db.Preload("User").First(&m, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (r *repository) GetByUserID(userID string) (*models.Merchant, error) {
	var m models.Merchant
	err := r.db.Preload("User").First(&m, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (r *repository) Update(merchant *models.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *repository) List() ([]models.Merchant, error) {
	var merchants []models.Merchant
	err := r.db.Preload("User").Order("created_at desc").Find(&merchants).Error
	return merchants, err
}
