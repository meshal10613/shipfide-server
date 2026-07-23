package rider

import (
	"net/http"

	"server/internal/domain/rider/dto"
	"server/internal/models"
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

func (h *Handler) CreateRider(c fiber.Ctx) error {
	var req dto.CreateRiderRequest
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

	userID, _ := c.Locals("user_id").(string)
	if req.UserID != "" {
		role, _ := c.Locals("role").(string)
		if role == string(models.RoleAdmin) || role == string(models.RoleSuperAdmin) {
			userID = req.UserID
		}
	}

	res, err := h.service.CreateRider(userID, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Rider profile created successfully",
		Data:    res,
	})
}

func (h *Handler) GetMyRiderProfile(c fiber.Ctx) error {
	userID, _ := c.Locals("user_id").(string)
	res, err := h.service.GetRiderByUserID(userID)
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

func (h *Handler) GetRiderByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetRiderByID(id)
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

func (h *Handler) UpdateRider(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateRiderRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.UpdateRider(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Rider profile updated successfully",
		Data:    res,
	})
}

func (h *Handler) UpdateKycStatus(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateKycStatusRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.UpdateKycStatus(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Rider KYC status updated successfully",
		Data:    res,
	})
}

func (h *Handler) UpdateStatus(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateRiderStatusRequest
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
		Message: "Rider active status updated successfully",
		Data:    res,
	})
}

func (h *Handler) ListRiders(c fiber.Ctx) error {
	res, err := h.service.ListRiders()
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
