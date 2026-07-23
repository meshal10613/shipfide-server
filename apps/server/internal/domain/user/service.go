package user

import (
	"context"
	"errors"

	"server/internal/config"
	"server/internal/domain/user/dto"
	"server/internal/models"
	cldService "server/pkg/cloudinary"
	querybuilder "server/pkg/queryBuilder"

	"github.com/gofiber/fiber/v3"
)

type Service interface {
	GetUsers(c fiber.Ctx, params querybuilder.QueryParams) ([]*userDto.UserResponse, int64, error)
	GetUser(id string) (*userDto.UserResponse, error)
	UpdateUser(id string, req *userDto.UpdateUserRequest, callerID string, callerRole string) (*userDto.UserResponse, error)
	DeleteUser(id string, callerID string, callerRole string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUsers(c fiber.Ctx, params querybuilder.QueryParams) ([]*userDto.UserResponse, int64, error) {
	users, total, err := s.repo.FindAll(c, params)
	if err != nil {
		return nil, 0, err
	}
	return userDto.MapToUserResponseList(users), total, nil
}

func (s *service) GetUser(id string) (*userDto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return userDto.MapToUserResponse(user), nil
}

func (s *service) UpdateUser(id string, req *userDto.UpdateUserRequest, callerID string, callerRole string) (*userDto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	isSelf := callerID == id
	isAdmin := callerRole == string(models.RoleAdmin)
	isSuperAdmin := callerRole == string(models.RoleSuperAdmin)

	// 1. Profile fields (name, phone, image) can ONLY be updated by the user themselves
	if req.Name != nil || req.Phone != nil || req.Image != nil {
		if !isSelf {
			return nil, errors.New("forbidden: you can only update your own name, phone, and image")
		}
	}

	// 2. Role and Status updates
	if req.Role != nil || req.Status != nil {
		if !isSuperAdmin && !isAdmin {
			return nil, errors.New("forbidden: you do not have permission to update user role or status")
		}

		if isAdmin {
			// Admin can update only merchant and rider role and status
			if user.Role != models.RoleMerchant && user.Role != models.RoleRider {
				return nil, errors.New("forbidden: admin can only update role and status of merchants and riders")
			}

			// If changing role, admin can only set it to merchant or rider
			if req.Role != nil && *req.Role != models.RoleMerchant && *req.Role != models.RoleRider {
				return nil, errors.New("forbidden: admin can only assign merchant or rider roles")
			}
		}
	}

	var oldImageURL string
	if user.Image != nil {
		oldImageURL = *user.Image
	}

	// Update fields if provided
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Image != nil {
		user.Image = req.Image
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	// Clean up previous image from Cloudinary if successfully replaced
	if req.Image != nil && oldImageURL != "" && oldImageURL != *req.Image {
		cld, _ := cldService.NewCloudinaryService(config.AppConfig)
		if cld != nil {
			_ = cld.DeleteImage(context.Background(), oldImageURL)
		}
	}

	return userDto.MapToUserResponse(user), nil
}

func (s *service) DeleteUser(id string, callerID string, callerRole string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	isSelf := callerID == id
	isAdmin := callerRole == string(models.RoleAdmin)
	isSuperAdmin := callerRole == string(models.RoleSuperAdmin)

	// Super admin cannot delete themselves
	if user.Role == models.RoleSuperAdmin && isSelf {
		return errors.New("forbidden: super admin cannot delete themselves")
	}

	// Clean up user avatar image from Cloudinary on delete
	if user.Image != nil && *user.Image != "" {
		cld, _ := cldService.NewCloudinaryService(config.AppConfig)
		if cld != nil {
			_ = cld.DeleteImage(context.Background(), *user.Image)
		}
	}

	// Super admin can delete anyone else
	if isSuperAdmin {
		return s.repo.Delete(id)
	}

	// Admin can delete themselves or any user who is a Merchant or Rider
	if isAdmin {
		if isSelf || user.Role == models.RoleMerchant || user.Role == models.RoleRider {
			return s.repo.Delete(id)
		}
		return errors.New("forbidden: admin can only delete merchants, riders, or themselves")
	}

	// Merchants and Riders can only delete themselves
	if isSelf {
		return s.repo.Delete(id)
	}

	return errors.New("forbidden: you do not have permission to delete this user")
}
