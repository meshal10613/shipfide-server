package merchant

import (
	"net/http"

	"server/internal/domain/merchant/dto"
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

func (h *Handler) CreateMerchant(c fiber.Ctx) error {
	var req dto.CreateMerchantRequest
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

	res, err := h.service.CreateMerchant(userID, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Merchant profile created successfully",
		Data:    res,
	})
}

func (h *Handler) GetMyMerchantProfile(c fiber.Ctx) error {
	userID, _ := c.Locals("user_id").(string)
	res, err := h.service.GetMerchantByUserID(userID)
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

func (h *Handler) GetMerchantByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetMerchantByID(id)
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

func (h *Handler) UpdateMerchant(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateMerchantRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.UpdateMerchant(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Merchant profile updated successfully",
		Data:    res,
	})
}

func (h *Handler) UpdateKyc(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateKycRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.UpdateKyc(id, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Merchant KYC updated successfully",
		Data:    res,
	})
}

func (h *Handler) ListMerchants(c fiber.Ctx) error {
	res, err := h.service.ListMerchants()
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
