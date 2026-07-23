package cloudinary

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"server/internal/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Service interface {
	UploadImage(ctx context.Context, file interface{}, folder string, oldImageURL string) (string, error)
	DeleteImage(ctx context.Context, imageURL string) error
	ExtractPublicID(imageURL string) string
}

type service struct {
	cld *cloudinary.Cloudinary
}

// NewCloudinaryService initializes a Cloudinary client strictly using CloudName, ApiKey, and ApiSecret.
func NewCloudinaryService(cfg *config.Config) (Service, error) {
	var cld *cloudinary.Cloudinary
	var err error

	if cfg != nil && cfg.CloudinaryCloudName != "" && cfg.CloudinaryApiKey != "" && cfg.CloudinaryApiSecret != "" {
		cld, err = cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryApiKey, cfg.CloudinaryApiSecret)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary client: %w", err)
	}

	return &service{cld: cld}, nil
}

// ExtractPublicID extracts the Cloudinary Public ID from a full image URL.
func ExtractPublicIDFromURL(imageURL string) string {
	if imageURL == "" {
		return ""
	}
	parts := strings.Split(imageURL, "/upload/")
	if len(parts) < 2 {
		return ""
	}
	path := parts[1]

	// Strip version prefix if present (e.g. v1678900000/)
	reVersion := regexp.MustCompile(`^v\d+/`)
	path = reVersion.ReplaceAllString(path, "")

	// Strip extension (.jpg, .png, etc.)
	ext := filepath.Ext(path)
	if ext != "" {
		path = strings.TrimSuffix(path, ext)
	}
	return path
}

func (s *service) ExtractPublicID(imageURL string) string {
	return ExtractPublicIDFromURL(imageURL)
}

// UploadImage uploads a file to Cloudinary under the specified folder.
// If oldImageURL is provided and non-empty, after successful upload of the new image,
// the previous image is destroyed from Cloudinary.
func (s *service) UploadImage(ctx context.Context, file interface{}, folder string, oldImageURL string) (string, error) {
	if s.cld == nil {
		return "", fmt.Errorf("Cloudinary is not configured on this server. Please provide CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, and CLOUDINARY_API_SECRET in environment")
	}

	uploadParams := uploader.UploadParams{
		Folder: folder,
	}

	resp, err := s.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Cloudinary: %w", err)
	}

	newImageURL := resp.SecureURL

	// If a previous image existed, delete it AFTER successful upload of the new image
	if oldImageURL != "" {
		oldPublicID := s.ExtractPublicID(oldImageURL)
		if oldPublicID != "" {
			// Asynchronously or inline attempt to delete previous image
			_, _ = s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
				PublicID: oldPublicID,
			})
		}
	}

	return newImageURL, nil
}

// DeleteImage destroys an image from Cloudinary by its URL.
func (s *service) DeleteImage(ctx context.Context, imageURL string) error {
	if s.cld == nil || imageURL == "" {
		return nil
	}

	publicID := s.ExtractPublicID(imageURL)
	if publicID == "" {
		return fmt.Errorf("invalid Cloudinary image URL: %s", imageURL)
	}

	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete image from Cloudinary: %w", err)
	}

	return nil
}
