package receiverFraud

import (
	"net/http"

	"server/internal/domain/receiverFraud/dto"
	httpresponse "server/pkg/httpResponse"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CheckPhone(c fiber.Ctx) error {
	phone := c.Query("phone")
	if phone == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "phone query parameter is required",
		})
	}

	res, err := h.service.CheckPhone(phone)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Data:    res,
	})
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetByID(id)
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

func (h *Handler) UpdateCodBlocked(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateCodBlockedRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
		})
	}

	res, err := h.service.UpdateCodBlocked(id, req.CODBlocked)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "COD block status updated",
		Data:    res,
	})
}

func (h *Handler) ListProfiles(c fiber.Ctx) error {
	res, err := h.service.ListProfiles()
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
