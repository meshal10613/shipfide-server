package session

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	httpresponse "server/pkg/httpResponse"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetSessions(c fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	var userIDStr string
	if u, ok := userIDVal.(uuid.UUID); ok {
		userIDStr = u.String()
	} else if s, ok := userIDVal.(string); ok {
		userIDStr = s
	}

	sessions, err := h.service.GetUserSessions(userIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to retrieve active sessions",
			Details: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "active sessions retrieved successfully",
		Data:    sessions,
	})
}

func (h *Handler) DeleteSession(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "session ID is required",
		})
	}

	userIDVal := c.Locals("user_id")
	var userIDStr string
	if u, ok := userIDVal.(uuid.UUID); ok {
		userIDStr = u.String()
	} else if s, ok := userIDVal.(string); ok {
		userIDStr = s
	}

	err := h.service.DeleteSession(id, userIDStr)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "session not found" {
			statusCode = http.StatusNotFound
		} else if len(err.Error()) >= 9 && err.Error()[:9] == "forbidden" {
			statusCode = http.StatusForbidden
		}
		return c.Status(statusCode).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "session deleted successfully",
	})
}
