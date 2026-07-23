package user

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"server/internal/domain/user/dto"
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

func (h *Handler) GetUsers(c fiber.Ctx) error {
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

	users, total, err := h.service.GetUsers(c, params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to retrieve users",
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
		Message: "users retrieved successfully",
		Data:    users,
		Meta:    meta,
	})
}

func (h *Handler) GetUserByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "user ID is required",
		})
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "user retrieved successfully",
		Data:    user,
	})
}

func (h *Handler) UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "user ID is required",
		})
	}

	var req userDto.UpdateUserRequest
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

	// Retrieve caller ID and role
	userIDVal := c.Locals("user_id")
	var callerID string
	if u, ok := userIDVal.(uuid.UUID); ok {
		callerID = u.String()
	} else if s, ok := userIDVal.(string); ok {
		callerID = s
	}

	roleVal := c.Locals("role")
	callerRole, _ := roleVal.(string)

	updatedUser, err := h.service.UpdateUser(id, &req, callerID, callerRole)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "user not found" {
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
		Message: "user updated successfully",
		Data:    updatedUser,
	})
}

func (h *Handler) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "user ID is required",
		})
	}

	// Retrieve caller ID and role
	userIDVal := c.Locals("user_id")
	var callerID string
	if u, ok := userIDVal.(uuid.UUID); ok {
		callerID = u.String()
	} else if s, ok := userIDVal.(string); ok {
		callerID = s
	}

	roleVal := c.Locals("role")
	callerRole, _ := roleVal.(string)

	err := h.service.DeleteUser(id, callerID, callerRole)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "user not found" {
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
		Message: "user deleted successfully",
	})
}
