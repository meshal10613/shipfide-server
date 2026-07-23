package admin

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"server/internal/domain/admin/dto"
	httpresponse "server/pkg/httpResponse"
	querybuilder "server/pkg/queryBuilder"
	"server/pkg/validation"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateAdmin(c fiber.Ctx) error {
	var req adminDto.CreateAdminRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
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

	admin, err := h.service.CreateAdmin(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to create admin",
			Details: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "admin created successfully",
		Data:    admin,
	})
}

func (h *Handler) GetAdmins(c fiber.Ctx) error {
	params := querybuilder.NewQueryParams()

	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			params.Page = page
		}
	}
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			params.Limit = limit
		}
	}
	params.Search = c.Query("search")
	params.SortBy = c.Query("sortBy")
	params.Order = c.Query("order")

	admins, total, err := h.service.GetAdmins(c, params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to retrieve admins",
			Details: err.Error(),
		})
	}

	totalPages := int64(math.Ceil(float64(total) / float64(params.Limit)))
	meta := &httpresponse.Meta{
		Page:      params.Page,
		Limit:     params.Limit,
		Total:     total,
		TotalPage: totalPages,
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "admins retrieved successfully",
		Data:    admins,
		Meta:    meta,
	})
}

func (h *Handler) GetAdminByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "admin ID is required",
		})
	}

	admin, err := h.service.GetAdmin(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "admin retrieved successfully",
		Data:    admin,
	})
}

func (h *Handler) UpdateAdmin(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "admin ID is required",
		})
	}

	var req adminDto.UpdateAdminRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
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

	updatedAdmin, err := h.service.UpdateAdmin(id, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "admin not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "admin updated successfully",
		Data:    updatedAdmin,
	})
}

func (h *Handler) DeleteAdmin(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "admin ID is required",
		})
	}

	err := h.service.DeleteAdmin(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "admin not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "admin deleted successfully",
	})
}
