package upload

import (
	"server/internal/config"
	"server/pkg/middlewares"
	"server/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func UploadRoutes(
	api fiber.Router,
	cfg *config.Config,
	jwt utils.JwtService,
) {
	handler := NewHandler(cfg)

	router := api.Group("/upload")

	router.Use(middlewares.Authentication(jwt))

	router.Post("/image", handler.UploadImage)
	router.Delete("/image", handler.DeleteImage)
}
