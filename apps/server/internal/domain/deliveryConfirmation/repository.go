package deliveryConfirmation

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetOtpByShipmentID(shipmentID string) (*models.DeliveryConfirmation, error)
	CreateOtp(dc *models.DeliveryConfirmation) error
	UpdateOtp(dc *models.DeliveryConfirmation) error

	GetCodByShipmentID(shipmentID string) (*models.CodDeliveryConfirmation, error)
	GetCodByID(id string) (*models.CodDeliveryConfirmation, error)
	CreateCod(cod *models.CodDeliveryConfirmation) error
	UpdateCod(cod *models.CodDeliveryConfirmation) error

	DB() *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) DB() *gorm.DB {
	return r.db
}

func (r *repository) GetOtpByShipmentID(shipmentID string) (*models.DeliveryConfirmation, error) {
	var dc models.DeliveryConfirmation
	err := r.db.First(&dc, "shipment_id = ?", shipmentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &dc, nil
}

func (r *repository) CreateOtp(dc *models.DeliveryConfirmation) error {
	return r.db.Create(dc).Error
}

func (r *repository) UpdateOtp(dc *models.DeliveryConfirmation) error {
	return r.db.Save(dc).Error
}

func (r *repository) GetCodByShipmentID(shipmentID string) (*models.CodDeliveryConfirmation, error) {
	var cod models.CodDeliveryConfirmation
	err := r.db.First(&cod, "shipment_id = ?", shipmentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cod, nil
}

func (r *repository) GetCodByID(id string) (*models.CodDeliveryConfirmation, error) {
	var cod models.CodDeliveryConfirmation
	err := r.db.First(&cod, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cod, nil
}

func (r *repository) CreateCod(cod *models.CodDeliveryConfirmation) error {
	return r.db.Create(cod).Error
}

func (r *repository) UpdateCod(cod *models.CodDeliveryConfirmation) error {
	return r.db.Save(cod).Error
}
