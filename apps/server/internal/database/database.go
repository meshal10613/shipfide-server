package database

import (
	"fmt"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/internal/config"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	} else {
		color.Green("🚀 Database connection successfully established!")
	}

	DB = db
	return db
}
