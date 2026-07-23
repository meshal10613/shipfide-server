package hub

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"server/internal/models"
	querybuilder "server/pkg/queryBuilder"
)

type Repository interface {
	Create(hub *models.Hub) error
	FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Hub, int64, error)
	FindByID(id string) (*models.Hub, error)
	Update(hub *models.Hub) error
	Delete(id string) error
	HasStaff(hubID string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(hub *models.Hub) error {
	return r.db.Create(hub).Error
}

func (r *repository) FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Hub, int64, error) {
	var hubs []models.Hub
	var total int64

	// Clone GORM DB for count query
	countDb := r.db.Model(&models.Hub{})
	builderForCount := querybuilder.New(countDb).Filter(c)
	if params.Search != "" {
		builderForCount = builderForCount.Search(params.Search, "name", "division", "district", "postal_code", "address")
	}
	if err := builderForCount.DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Main query
	dbInstance := r.db.Model(&models.Hub{})
	builder := querybuilder.New(dbInstance).
		Filter(c).
		Search(params.Search, "name", "division", "district", "postal_code", "address").
		Sort(params.SortBy, params.Order).
		Paginate(params.Page, params.Limit)

	if err := builder.DB.Find(&hubs).Error; err != nil {
		return nil, 0, err
	}

	return hubs, total, nil
}

func (r *repository) FindByID(id string) (*models.Hub, error) {
	var hub models.Hub
	if err := r.db.Where("id = ?", id).First(&hub).Error; err != nil {
		return nil, err
	}
	return &hub, nil
}

func (r *repository) Update(hub *models.Hub) error {
	return r.db.Save(hub).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Hub{}, "id = ?", id).Error
}

func (r *repository) HasStaff(hubID string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Admin{}).Where("hub_id = ?", hubID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
