package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/satioO/iam/internal/entities"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/pkg/log"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]dtos.ListUserDTO, error)
	GetUserDetails(ctx context.Context) (*dtos.GetUserDetailsDTO, error)
	CreateUser(ctx context.Context, user *dtos.CreateUserDTO) error
	UpdateUser(ctx context.Context, userId uuid.UUID, user *dtos.UpdateUserDTO) error
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}

type userUsecase struct {
	db     *gorm.DB
	logger log.Factory
}

func NewUserUsecase(db *gorm.DB, logger log.Factory) *userUsecase {
	return &userUsecase{db, logger}
}

func (u userUsecase) GetUsers(ctx context.Context) ([]dtos.ListUserDTO, error) {
	u.logger.For(ctx).Info("Fetching users:::")

	var users []entities.User

	if err := u.db.Preload("Attributes").Find(&users).Error; err != nil {
		return nil, err
	}

	result := []dtos.ListUserDTO{}

	for _, user := range users {
		result = append(result, dtos.ListUserDTO{
			ID:                  user.ID,
			FirstName:           user.FirstName,
			LastName:            user.LastName,
			Username:            user.Username,
			Email:               user.Email,
			EmailVerified:       *user.EmailVerified,
			PhoneCountryCode:    user.PhoneCountryCode,
			PhoneNumber:         user.PhoneNumber,
			PhoneNumberVerified: *user.PhoneNumberVerified,
			Status:              user.Status,
		})
	}

	return result, nil
}

func (u userUsecase) GetUserDetails(ctx context.Context) (*dtos.GetUserDetailsDTO, error) {
	return nil, errors.New("Not Implemented")
}

func (u userUsecase) CreateUser(ctx context.Context, user *dtos.CreateUserDTO) error {
	return errors.New("Not Implemented")
}

func (u userUsecase) UpdateUser(ctx context.Context, userId uuid.UUID, user *dtos.UpdateUserDTO) error {
	return errors.New("Not Implemented")
}

func (u userUsecase) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	return errors.New("Not Implemented")
}
