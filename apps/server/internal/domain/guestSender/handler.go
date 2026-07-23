package guestSender

import (
	"net/http"

	"server/internal/domain/guestSender/dto"
	httpresponse "server/pkg/httpResponse"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateGuestSender(c fiber.Ctx) error {
	var req dto.CreateGuestSenderRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	adminID, _ := c.Locals("user_id").(string)

	res, err := h.service.CreateGuestSender(adminID, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Guest sender registered successfully",
		Data:    res,
	})
}

func (h *Handler) GetGuestSenderByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetGuestSenderByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Data:    res,
	})
}

func (h *Handler) FlagGuestSender(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.FlagGuestSenderRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.FlagGuestSender(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Guest sender flag status updated",
		Data:    res,
	})
}

func (h *Handler) ListGuestSenders(c fiber.Ctx) error {
	res, err := h.service.ListGuestSenders()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Data:    res,
	})
}
