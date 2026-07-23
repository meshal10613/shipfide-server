package guestSender

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(guest *models.GuestSender) error
	GetByID(id string) (*models.GuestSender, error)
	GetByPhone(phone string) (*models.GuestSender, error)
	Update(guest *models.GuestSender) error
	List() ([]models.GuestSender, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(guest *models.GuestSender) error {
	return r.db.Create(guest).Error
}

func (r *repository) GetByID(id string) (*models.GuestSender, error) {
	var g models.GuestSender
	err := r.db.Preload("Admin").First(&g, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *repository) GetByPhone(phone string) (*models.GuestSender, error) {
	var g models.GuestSender
	err := r.db.Preload("Admin").First(&g, "phone = ?", phone).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *repository) Update(guest *models.GuestSender) error {
	return r.db.Save(guest).Error
}

func (r *repository) List() ([]models.GuestSender, error) {
	var guests []models.GuestSender
	err := r.db.Preload("Admin").Order("created_at desc").Find(&guests).Error
	return guests, err
}
