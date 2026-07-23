package withdrawal

import (
	"net/http"

	"server/internal/domain/withdrawal/dto"
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

func (h *Handler) CreateWithdrawal(c fiber.Ctx) error {
	var req dto.CreateWithdrawalRequest
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

	res, err := h.service.CreateWithdrawal(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Withdrawal request submitted successfully",
		Data:    res,
	})
}

func (h *Handler) GetWithdrawalByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetWithdrawalByID(id)
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

func (h *Handler) UpdateStatus(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateWithdrawalStatusRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.UpdateStatus(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Withdrawal status updated successfully",
		Data:    res,
	})
}

func (h *Handler) ListWithdrawals(c fiber.Ctx) error {
	var riderID *string
	var merchantID *string

	if r := c.Query("riderId"); r != "" {
		riderID = &r
	}
	if m := c.Query("merchantId"); m != "" {
		merchantID = &m
	}

	res, err := h.service.ListWithdrawals(riderID, merchantID)
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
