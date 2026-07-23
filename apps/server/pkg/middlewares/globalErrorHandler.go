package middlewares

import (
	"errors"
	"server/pkg/httpResponse"

	"github.com/gofiber/fiber/v3"
)

// GlobalErrorHandler handles all errors returned from routes and middlewares
func GlobalErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(httpResponse.Error{
		Success: false,
		Message: err.Error(),
	})
}
