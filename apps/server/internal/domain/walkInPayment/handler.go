package walkInPayment

import (
	"net/http"

	"server/internal/domain/walkInPayment/dto"
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

func (h *Handler) CreatePayment(c fiber.Ctx) error {
	var req dto.CreateWalkInPaymentRequest
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

	res, err := h.service.CreatePayment(adminID, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Walk-in payment recorded successfully",
		Data:    res,
	})
}

func (h *Handler) ListPayments(c fiber.Ctx) error {
	res, err := h.service.ListPayments()
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
