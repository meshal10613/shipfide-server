package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"server/internal/models"
	querybuilder "server/pkg/queryBuilder"
)

type Repository interface {
	FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.User, int64, error)
	FindByID(id string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Clone GORM DB for count query
	countDb := r.db.Model(&models.User{})
	builderForCount := querybuilder.New(countDb).Filter(c)
	if params.Search != "" {
		builderForCount = builderForCount.Search(params.Search, "name", "email", "phone")
	}
	if err := builderForCount.DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Main query
	dbInstance := r.db.Model(&models.User{})
	builder := querybuilder.New(dbInstance).
		Filter(c).
		Search(params.Search, "name", "email", "phone").
		Sort(params.SortBy, params.Order).
		Paginate(params.Page, params.Limit)

	if err := builder.DB.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *repository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
