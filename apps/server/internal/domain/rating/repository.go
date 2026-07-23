package rating

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateDeliveryRating(rating *models.DeliveryRating) error
	ListDeliveryRatingsByRiderID(riderID string) ([]models.DeliveryRating, error)

	CreateMerchantRating(rating *models.MerchantDeliveryRating) error
	ListMerchantRatings() ([]models.MerchantDeliveryRating, error)

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

func (r *repository) CreateDeliveryRating(rating *models.DeliveryRating) error {
	return r.db.Create(rating).Error
}

func (r *repository) ListDeliveryRatingsByRiderID(riderID string) ([]models.DeliveryRating, error) {
	var ratings []models.DeliveryRating
	err := r.db.Where("rider_id = ?", riderID).Order("created_at desc").Find(&ratings).Error
	return ratings, err
}

func (r *repository) CreateMerchantRating(rating *models.MerchantDeliveryRating) error {
	return r.db.Create(rating).Error
}

func (r *repository) ListMerchantRatings() ([]models.MerchantDeliveryRating, error) {
	var ratings []models.MerchantDeliveryRating
	err := r.db.Preload("Merchant").Preload("Shipment").Order("created_at desc").Find(&ratings).Error
	return ratings, err
}
