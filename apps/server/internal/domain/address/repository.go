package address

import (
	"errors"

	"server/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(address *models.Address) error
	GetByID(id string) (*models.Address, error)
	Update(address *models.Address) error
	List() ([]models.Address, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *repository) GetByID(id string) (*models.Address, error) {
	var address models.Address
	err := r.db.First(&address, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}

func (r *repository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *repository) List() ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Order("created_at desc").Find(&addresses).Error
	return addresses, err
}
