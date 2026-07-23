package shipment

import (
	"net/http"
	"strconv"

	"server/internal/domain/shipment/dto"
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

func (h *Handler) CalculatePrice(c fiber.Ctx) error {
	var req dto.CalculatePriceRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.CalculatePrice(req)
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

func (h *Handler) CreateShipment(c fiber.Ctx) error {
	var req dto.CreateShipmentRequest
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

	var adminID *string
	userID, _ := c.Locals("user_id").(string)
	role, _ := c.Locals("role").(string)
	if role == string(models.RoleAdmin) || role == string(models.RoleSuperAdmin) {
		adminID = &userID
	}

	res, err := h.service.CreateShipment(adminID, req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Shipment created successfully",
		Data:    res,
	})
}

func (h *Handler) TrackShipment(c fiber.Ctx) error {
	code := c.Params("trackingCode")
	res, err := h.service.GetShipmentByTrackingCode(code)
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

func (h *Handler) GetShipmentByID(c fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.service.GetShipmentByID(id)
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
	var req dto.UpdateShipmentStatusRequest
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
		Message: "Shipment status updated",
		Data:    res,
	})
}

func (h *Handler) AssignRider(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.AssignRiderRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.AssignRider(id, req.RiderID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Rider assigned successfully",
		Data:    res,
	})
}

func (h *Handler) AssignHub(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.AssignHubRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	res, err := h.service.AssignHub(id, req.HubID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Hub assigned successfully",
		Data:    res,
	})
}

func (h *Handler) ListShipments(c fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit

	filter := ListFilter{
		Status:     models.ShipmentStatus(c.Query("status")),
		SenderType: models.SenderType(c.Query("senderType")),
		MerchantID: c.Query("merchantId"),
		RiderID:    c.Query("riderId"),
		HubID:      c.Query("hubId"),
		Search:     c.Query("search"),
		Limit:      limit,
		Offset:     offset,
	}

	shipments, total, err := h.service.ListShipments(filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Data: fiber.Map{
			"shipments": shipments,
			"total":     total,
			"page":      page,
			"limit":     limit,
		},
	})
}
