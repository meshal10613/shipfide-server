package withdrawal

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(withdrawal *models.Withdrawal) error
	GetByID(id string) (*models.Withdrawal, error)
	Update(withdrawal *models.Withdrawal) error
	List(riderID *string, merchantID *string) ([]models.Withdrawal, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(withdrawal *models.Withdrawal) error {
	return r.db.Create(withdrawal).Error
}

func (r *repository) GetByID(id string) (*models.Withdrawal, error) {
	var w models.Withdrawal
	err := r.db.Preload("Rider").Preload("Merchant").First(&w, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &w, nil
}

func (r *repository) Update(withdrawal *models.Withdrawal) error {
	return r.db.Save(withdrawal).Error
}

func (r *repository) List(riderID *string, merchantID *string) ([]models.Withdrawal, error) {
	var withdrawals []models.Withdrawal
	q := r.db.Preload("Rider").Preload("Merchant")

	if riderID != nil {
		q = q.Where("rider_id = ?", *riderID)
	}
	if merchantID != nil {
		q = q.Where("merchant_id = ?", *merchantID)
	}

	err := q.Order("requested_at desc").Find(&withdrawals).Error
	return withdrawals, err
}
