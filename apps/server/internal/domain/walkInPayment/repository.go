package walkInPayment

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(payment *models.WalkInPayment) error
	GetByID(id string) (*models.WalkInPayment, error)
	List() ([]models.WalkInPayment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(payment *models.WalkInPayment) error {
	return r.db.Create(payment).Error
}

func (r *repository) GetByID(id string) (*models.WalkInPayment, error) {
	var p models.WalkInPayment
	err := r.db.Preload("CollectedByAdmin").Preload("Shipment").First(&p, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *repository) List() ([]models.WalkInPayment, error) {
	var payments []models.WalkInPayment
	err := r.db.Preload("CollectedByAdmin").Preload("Shipment").Order("created_at desc").Find(&payments).Error
	return payments, err
}
