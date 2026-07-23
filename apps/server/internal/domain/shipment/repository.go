package shipment

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type ListFilter struct {
	Status     models.ShipmentStatus
	SenderType models.SenderType
	MerchantID string
	RiderID    string
	HubID      string
	Search     string
	Limit      int
	Offset     int
}

type Repository interface {
	Create(shipment *models.Shipment) error
	GetByID(id string) (*models.Shipment, error)
	GetByTrackingCode(trackingCode string) (*models.Shipment, error)
	Update(shipment *models.Shipment) error
	List(filter ListFilter) ([]models.Shipment, int64, error)
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

func (r *repository) Create(shipment *models.Shipment) error {
	return r.db.Create(shipment).Error
}

func (r *repository) GetByID(id string) (*models.Shipment, error) {
	var s models.Shipment
	err := r.db.Preload("Merchant").Preload("GuestSender").Preload("Rider").Preload("Hub").First(&s, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *repository) GetByTrackingCode(trackingCode string) (*models.Shipment, error) {
	var s models.Shipment
	err := r.db.Preload("Merchant").Preload("GuestSender").Preload("Rider").Preload("Hub").First(&s, "tracking_code = ?", trackingCode).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *repository) Update(shipment *models.Shipment) error {
	return r.db.Save(shipment).Error
}

func (r *repository) List(filter ListFilter) ([]models.Shipment, int64, error) {
	var shipments []models.Shipment
	var total int64

	q := r.db.Model(&models.Shipment{})

	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}
	if filter.SenderType != "" {
		q = q.Where("sender_type = ?", filter.SenderType)
	}
	if filter.MerchantID != "" {
		q = q.Where("merchant_id = ?", filter.MerchantID)
	}
	if filter.RiderID != "" {
		q = q.Where("rider_id = ?", filter.RiderID)
	}
	if filter.HubID != "" {
		q = q.Where("hub_id = ?", filter.HubID)
	}
	if filter.Search != "" {
		term := "%" + filter.Search + "%"
		q = q.Where("tracking_code ILIKE ? OR receiver_name ILIKE ? OR receiver_phone ILIKE ? OR sender_name ILIKE ?", term, term, term, term)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		q = q.Offset(filter.Offset)
	}

	err := q.Preload("Merchant").Preload("GuestSender").Preload("Rider").Preload("Hub").Order("created_at desc").Find(&shipments).Error
	return shipments, total, err
}
