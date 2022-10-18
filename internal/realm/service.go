package realm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/satioO/iam/internal/entities"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/pkg/log"
	"gorm.io/gorm"
)

const (
	ACCESS_TOKEN_LIFESPAN  = "300"
	REFRESH_TOKEN_LIFESPAN = "600"
)

type RealmUsecase interface {
	GetRealms(ctx context.Context) ([]dtos.ListRealmDTO, error)
	GetRealmByID(realmId uuid.UUID) (*dtos.GetRealmDTO, error)
	CreateRealm(realm dtos.CreateRealmDTO) (*uuid.UUID, error)
	UpdateRealm(realmId uuid.UUID, realm dtos.UpdateRealmDTO) (bool, error)
	DeleteRealm(realmId uuid.UUID) (bool, error)
}

type usecase struct {
	db     *gorm.DB
	logger log.Factory
}

func NewRealmUsecase(db *gorm.DB, logger log.Factory) *usecase {
	return &usecase{
		db,
		logger,
	}
}

func (u usecase) GetRealms(ctx context.Context) ([]dtos.ListRealmDTO, error) {
	u.logger.For(ctx).Info("Fetching realms:::")

	var foundRealms []entities.Realm
	if err := u.db.Preload("Attributes").Where("enabled = ?", true).Find(&foundRealms).Error; err != nil {
		return nil, err
	}

	result := []dtos.ListRealmDTO{}
	for _, realm := range foundRealms {
		attributes := make(map[string]string)
		for _, attribute := range realm.Attributes {
			attributes[attribute.Name] = attribute.Value
		}

		result = append(result, dtos.ListRealmDTO{
			ID:                      realm.ID,
			Name:                    realm.Name,
			DisplayName:             realm.DisplayName,
			Logo:                    realm.Logo,
			SupportURL:              realm.SupportURL,
			SupportEmail:            realm.SupportEmail,
			DuplicateEmailAllowed:   realm.DuplicateEmailAllowed,
			DuplicatePhoneAllowed:   realm.DuplicatePhoneAllowed,
			RegisterEmailAsUsername: realm.RegisterEmailAsUsername,
			RegisterPhoneAsUsername: realm.RegisterPhoneAsUsername,
			Enabled:                 realm.Enabled,
			Attributes:              attributes,
		})
	}

	return result, nil
}

func (u usecase) GetRealmByID(realmId uuid.UUID) (*dtos.GetRealmDTO, error) {
	var foundRealm entities.Realm

	if err := u.db.Where("id = ?", realmId).Preload("Attributes").First(&foundRealm).Error; err != nil {
		return nil, err
	}

	attributes := make(map[string]string)
	for _, attribute := range foundRealm.Attributes {
		attributes[attribute.Name] = attribute.Value
	}

	return &dtos.GetRealmDTO{
		ID:                      foundRealm.ID,
		Name:                    foundRealm.Name,
		DisplayName:             foundRealm.DisplayName,
		Logo:                    foundRealm.Logo,
		SupportURL:              foundRealm.SupportURL,
		SupportEmail:            foundRealm.SupportEmail,
		DuplicateEmailAllowed:   foundRealm.DuplicateEmailAllowed,
		DuplicatePhoneAllowed:   foundRealm.DuplicatePhoneAllowed,
		RegisterEmailAsUsername: foundRealm.RegisterEmailAsUsername,
		RegisterPhoneAsUsername: foundRealm.RegisterPhoneAsUsername,
		Enabled:                 foundRealm.Enabled,
		Attributes:              attributes,
	}, nil
}

func (u usecase) CreateRealm(body dtos.CreateRealmDTO) (*uuid.UUID, error) {
	u.logger.Bg().Info(fmt.Sprintf("Creating realm::: %s", body.Name))
	defaultAttributes := map[string]string{
		"access_token_lifespan":  ACCESS_TOKEN_LIFESPAN,
		"refresh_token_lifespan": REFRESH_TOKEN_LIFESPAN,
	}

	for name, value := range body.Attributes {
		defaultAttributes[name] = value
	}

	var attributes []entities.RealmAttribute
	for name, value := range defaultAttributes {
		attributes = append(attributes, entities.RealmAttribute{
			Name:  name,
			Value: value,
		})
	}

	createdRealm := &entities.Realm{
		Name:                    body.Name,
		DisplayName:             body.DisplayName,
		Logo:                    body.Logo,
		SupportURL:              body.SupportURL,
		SupportEmail:            body.SupportEmail,
		DuplicateEmailAllowed:   body.DuplicateEmailAllowed,
		DuplicatePhoneAllowed:   body.DuplicatePhoneAllowed,
		RegisterEmailAsUsername: body.RegisterEmailAsUsername,
		RegisterPhoneAsUsername: body.RegisterPhoneAsUsername,
		Enabled:                 body.Enabled,
		Attributes:              attributes,
	}

	if err := u.db.Create(&createdRealm).Error; err != nil {
		return nil, err
	}

	return &createdRealm.ID, nil
}

func (u usecase) UpdateRealm(realmId uuid.UUID, realm dtos.UpdateRealmDTO) (bool, error) {
	u.logger.Bg().Info(fmt.Sprintf("Updating realm::: %s", realmId))

	updatedRealm := entities.Realm{
		Name:                    realm.Name,
		DisplayName:             realm.DisplayName,
		Logo:                    realm.Logo,
		SupportURL:              realm.SupportURL,
		SupportEmail:            realm.SupportEmail,
		DuplicateEmailAllowed:   realm.DuplicateEmailAllowed,
		DuplicatePhoneAllowed:   realm.DuplicatePhoneAllowed,
		RegisterEmailAsUsername: realm.RegisterEmailAsUsername,
		RegisterPhoneAsUsername: realm.DuplicatePhoneAllowed,
	}

	if err := u.db.Where("id = ?", realmId).Updates(&updatedRealm).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u usecase) DeleteRealm(realmId uuid.UUID) (bool, error) {
	status := false
	deletedRealm := entities.Realm{Enabled: &status}

	query := u.db.Where("id = ?", realmId).Updates(&deletedRealm)
	u.logger.Bg().Info(fmt.Sprintf("rows affected: %d", query.RowsAffected))

	if err := query.Error; err != nil {
		return false, err
	}

	return true, nil
}
