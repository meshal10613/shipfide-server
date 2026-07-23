package server

import (
	"log"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/internal/config"
	"server/internal/models"
	"server/internal/routes"
	"server/pkg/middlewares"
	"server/pkg/seed"
	"server/pkg/utils"
	"server/pkg/validation"
)

// StartServer initializes configurations, runs migrations, seeds data, and starts the Fiber HTTP server
func StartServer(db *gorm.DB, cfg *config.Config) {
	// Run database auto-migrations
	color.Yellow("🔄 Running database auto-migrations...")
	err := models.Migrate(db)
	if err != nil {
		color.Red("❌ Database auto-migration failed: %v", err)
		log.Fatalf("Migration error: %v", err)
	}
	color.Green("✅ Database auto-migrations completed successfully!")

	// Seed default super admin
	color.Yellow("⏳ Checking super admin seed...")
	seed.SeedSuperAdmin(db, seed.SuperAdminSeed{
		Name:     cfg.SuperAdminName,
		Email:    cfg.SuperAdminEmail,
		Password: cfg.SuperAdminPassword,
	})

	// Initialize custom validator
	valInstance := validation.NewCustomValidator(validator.New())

	// Initialize JWT service
	jwtService := utils.NewJwtService(cfg.JwtSecretKey)

	// Initialize Fiber application
	app := fiber.New(fiber.Config{
		AppName:      "Shipfide API Server v1.0",
		ErrorHandler: middlewares.GlobalErrorHandler,
	})

	// Setup all API Routes
	routes.SetupRoutes(app, db, valInstance, jwtService)

	// Custom 404 Not Found handler
	app.Use(middlewares.NotFoundHandler)

	// Start the server
	addr := ":" + cfg.Port
	color.Magenta("📡 Server starting on address %s...", addr)

	if err := app.Listen(addr); err != nil {
		color.Red("❌ Server failed to start: %v", err)
		log.Fatalf("Server error: %v", err)
	}
}
