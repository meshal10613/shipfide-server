package middlewares

import (
	"server/pkg/httpResponse"

	"github.com/gofiber/fiber/v3"
)

// NotFoundHandler handles unmatched routes (404 Not Found)
func NotFoundHandler(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(httpResponse.Error{
		Success: false,
		Message: "route not found",
		Details: "Cannot " + c.Method() + " " + c.Path(),
	})
}
