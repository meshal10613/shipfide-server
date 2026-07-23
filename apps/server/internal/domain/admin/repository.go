package admin

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"server/internal/models"
	querybuilder "server/pkg/queryBuilder"
)

type Repository interface {
	CreateAdmin(user *models.User, account *models.Account, admin *models.Admin) error
	FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Admin, int64, error)
	FindByID(id string) (*models.Admin, error)
	FindByUserID(userID string) (*models.Admin, error)
	Update(admin *models.Admin, user *models.User) error
	Delete(id string) error
	GetUserByEmail(email string) (*models.User, error)
	GetHubByID(hubID string) (*models.Hub, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateAdmin(user *models.User, account *models.Account, admin *models.Admin) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		account.UserID = user.ID
		if err := tx.Create(account).Error; err != nil {
			return err
		}
		admin.UserID = user.ID
		if err := tx.Create(admin).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Admin, int64, error) {
	var admins []models.Admin
	var total int64

	// Clone GORM DB for count query
	countDb := r.db.Model(&models.Admin{}).Joins("User")
	builderForCount := querybuilder.New(countDb).Filter(c)
	if params.Search != "" {
		builderForCount = builderForCount.Search(params.Search, "admins.name", "admins.phone", "User.email")
	}
	if err := builderForCount.DB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Main query
	dbInstance := r.db.Model(&models.Admin{}).Preload("User").Preload("Hub")
	builder := querybuilder.New(dbInstance).
		Filter(c).
		Search(params.Search, "admins.name", "admins.phone", "User.email").
		Sort(params.SortBy, params.Order).
		Paginate(params.Page, params.Limit)

	if err := builder.DB.Find(&admins).Error; err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

func (r *repository) FindByID(id string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.db.Preload("User").Preload("Hub").Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *repository) FindByUserID(userID string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.db.Preload("User").Preload("Hub").Where("user_id = ?", userID).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *repository) Update(admin *models.Admin, user *models.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if user != nil {
			if err := tx.Save(user).Error; err != nil {
				return err
			}
		}
		if err := tx.Save(admin).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) Delete(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var admin models.Admin
		if err := tx.Where("id = ?", id).First(&admin).Error; err != nil {
			return err
		}
		// Deleting the user will cascade delete the Admin profile because of CASCADE constraint.
		// However, doing it explicitly or deleting the User is cleaner.
		if err := tx.Delete(&models.User{}, "id = ?", admin.UserID).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetHubByID(hubID string) (*models.Hub, error) {
	var hub models.Hub
	if err := r.db.Where("id = ?", hubID).First(&hub).Error; err != nil {
		return nil, err
	}
	return &hub, nil
}
