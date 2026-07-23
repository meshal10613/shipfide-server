package hub

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"server/internal/domain/hub/dto"
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

func (h *Handler) CreateHub(c fiber.Ctx) error {
	var req hubDto.CreateHubRequest
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

	hub, err := h.service.CreateHub(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to create hub",
			Details: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "hub created successfully",
		Data:    hub,
	})
}

func (h *Handler) GetHubs(c fiber.Ctx) error {
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

	hubs, total, err := h.service.GetHubs(c, params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to retrieve hubs",
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
		Message: "hubs retrieved successfully",
		Data:    hubs,
		Meta:    meta,
	})
}

func (h *Handler) GetHubByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "hub ID is required",
		})
	}

	hub, err := h.service.GetHub(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "hub retrieved successfully",
		Data:    hub,
	})
}

func (h *Handler) UpdateHub(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "hub ID is required",
		})
	}

	var req hubDto.UpdateHubRequest
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

	updatedHub, err := h.service.UpdateHub(id, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "hub not found" {
			statusCode = http.StatusNotFound
		}
		return c.Status(statusCode).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "hub updated successfully",
		Data:    updatedHub,
	})
}

func (h *Handler) DeleteHub(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "hub ID is required",
		})
	}

	err := h.service.DeleteHub(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "hub not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "cannot delete hub: staff are still assigned. Reassign staff to another hub first" {
			statusCode = http.StatusBadRequest
		}
		return c.Status(statusCode).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "hub deleted successfully",
	})
}
