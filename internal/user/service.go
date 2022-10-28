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

	var foundUsers []entities.User

	if err := u.db.Preload("Attributes").Find(&foundUsers).Error; err != nil {
		return nil, err
	}

	result := []dtos.ListUserDTO{}

	for _, foundUser := range foundUsers {
		result = append(result, dtos.ListUserDTO{
			ID:                  foundUser.ID,
			FirstName:           foundUser.FirstName,
			LastName:            foundUser.LastName,
			Username:            foundUser.Username,
			Email:               foundUser.Email,
			EmailVerified:       *foundUser.EmailVerified,
			PhoneCountryCode:    foundUser.PhoneCountryCode,
			PhoneNumber:         foundUser.PhoneNumber,
			PhoneNumberVerified: *foundUser.PhoneNumberVerified,
			Status:              foundUser.Status.String(),
		})
	}

	return result, nil
}

func (u userUsecase) GetUserDetails(ctx context.Context) (*dtos.GetUserDetailsDTO, error) {
	u.logger.For(ctx).Info("Fetching User Details:::")

	var foundUser entities.User

	if err := u.db.Preload("Attributes").First(&foundUser).Error; err != nil {
		return nil, err
	}

	result := dtos.GetUserDetailsDTO{
		ID:                  foundUser.ID,
		FirstName:           foundUser.FirstName,
		LastName:            foundUser.LastName,
		Email:               foundUser.Email,
		EmailVerified:       *foundUser.EmailVerified,
		PhoneCountryCode:    foundUser.PhoneCountryCode,
		PhoneNumber:         foundUser.PhoneNumber,
		PhoneNumberVerified: *foundUser.PhoneNumberVerified,
		Status:              foundUser.Status.String(),
	}

	return &result, nil
}

func (u userUsecase) CreateUser(ctx context.Context, user *dtos.CreateUserDTO) error {
	u.logger.For(ctx).Info("Creating User:::")

	var status dtos.UserStatus = dtos.Active

	entity := entities.User{
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		Username:            user.Username,
		Email:               user.Email,
		EmailVerified:       &user.EmailVerified,
		PhoneCountryCode:    user.PhoneCountryCode,
		PhoneNumber:         user.PhoneNumber,
		PhoneNumberVerified: &user.PhoneNumberVerified,
		Status:              status,
	}

	if err := u.db.Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

func (u userUsecase) UpdateUser(ctx context.Context, userId uuid.UUID, user *dtos.UpdateUserDTO) error {
	return errors.New("Not Implemented")
}
