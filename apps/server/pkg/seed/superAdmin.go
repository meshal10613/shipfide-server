package seed

import (
	"errors"
	"server/internal/models"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SuperAdminSeed struct {
	Name     string
	Email    string
	Password string
}

func SeedSuperAdmin(db *gorm.DB, admin SuperAdminSeed) {
	var existing models.User

	err := db.Where("email = ?", admin.Email).First(&existing).Error

	// Super admin already exists — skip seeding
	if err == nil {
		color.Yellow("⚠️  Super admin already exists — skipping seed  [%s]", existing.Email)
		return
	}

	// Unexpected DB error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		color.Red("❌ Failed to check super admin existence: %v", err)
		return
	}

	userID := uuid.NewString()
	user := models.User{
		ID:     userID,
		Name:   admin.Name,
		Email:  admin.Email,
		Role:   models.RoleSuperAdmin,
		Status: models.UserStatusActive,
	}

	account := models.Account{
		ID:     uuid.NewString(),
		UserID: userID,
	}

	if err := account.HashPassword(admin.Password); err != nil {
		color.Red("❌ Failed to hash super admin password: %v", err)
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		if err := tx.Create(&account).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		color.Red("❌ Failed to seed super admin user: %v", err)
		return
	}

	color.Green("✅ Super admin seeded successfully  [%s]", user.Email)
}
