package upload

import (
	"context"
	"net/http"

	"server/internal/config"
	cldService "server/pkg/cloudinary"
	httpresponse "server/pkg/httpResponse"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	cloudinary cldService.Service
}

func NewHandler(cfg *config.Config) *Handler {
	cld, _ := cldService.NewCloudinaryService(cfg)
	return &Handler{cloudinary: cld}
}

func (h *Handler) UploadImage(c fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fileHeader, err = c.FormFile("image")
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "No file uploaded. Please attach a file named 'file' or 'image'",
			Details: err.Error(),
		})
	}

	folder := c.FormValue("folder", "shipfide")
	oldImageURL := c.FormValue("oldImageUrl", "")

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to open uploaded file",
			Details: err.Error(),
		})
	}
	defer file.Close()

	if h.cloudinary == nil {
		return c.Status(http.StatusServiceUnavailable).JSON(httpresponse.Error{
			Success: false,
			Message: "Cloudinary service is not configured. Please set CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, and CLOUDINARY_API_SECRET in environment.",
		})
	}

	url, err := h.cloudinary.UploadImage(context.Background(), file, folder, oldImageURL)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Image uploaded successfully",
		Data: fiber.Map{
			"url": url,
		},
	})
}

func (h *Handler) DeleteImage(c fiber.Ctx) error {
	imageURL := c.Query("url")
	if imageURL == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Query parameter 'url' is required",
		})
	}

	if h.cloudinary == nil {
		return c.Status(http.StatusServiceUnavailable).JSON(httpresponse.Error{
			Success: false,
			Message: "Cloudinary service is not configured",
		})
	}

	if err := h.cloudinary.DeleteImage(context.Background(), imageURL); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(httpresponse.Success{
		Success: true,
		Message: "Image deleted from Cloudinary",
	})
}
