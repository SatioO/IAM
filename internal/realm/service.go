package realm

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/satioO/iam/pkg/dtos"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RealmUsecase interface {
	GetRealms() []dtos.ListRealmDTO
	GetRealmByID(realmId uuid.UUID) (*dtos.GetRealmDTO, error)
	CreateRealm(realm dtos.CreateRealmDTO) (*uuid.UUID, error)
	UpdateRealm(realmId uuid.UUID, realm dtos.UpdateRealmDTO) (*uuid.UUID, error)
	DeleteRealm(realmId uuid.UUID) error
}

type usecase struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRealmUsecase(db *gorm.DB, logger *zap.Logger) *usecase {
	return &usecase{
		db,
		logger,
	}
}

func (r usecase) GetRealms() []dtos.ListRealmDTO {
	r.logger.Info("Fetching realms:::")
	var foundRealms []Realm
	r.db.Preload("Attributes").Where("enabled = ?", true).Find(&foundRealms)

	var result []dtos.ListRealmDTO
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

	return result
}

func (r usecase) GetRealmByID(realmId uuid.UUID) (*dtos.GetRealmDTO, error) {
	var foundRealm Realm

	if err := r.db.Where("id = ?", realmId).Preload("Attributes").First(&foundRealm).Error; err != nil {
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

func (r usecase) CreateRealm(body dtos.CreateRealmDTO) (*uuid.UUID, error) {
	r.logger.Info(fmt.Sprintf("Creating realm::: %s", body.Name))
	defaultAttributes := map[string]string{
		"access_token_lifespan":  ACCESS_TOKEN_LIFESPAN,
		"refresh_token_lifespan": REFRESH_TOKEN_LIFESPAN,
	}

	for name, value := range body.Attributes {
		defaultAttributes[name] = value
	}

	var attributes []RealmAttribute
	for name, value := range defaultAttributes {
		attributes = append(attributes, RealmAttribute{
			Name:  name,
			Value: value,
		})
	}

	createdRealm := &Realm{
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

	if err := r.db.Create(&createdRealm).Error; err != nil {
		return nil, err
	}

	return &createdRealm.ID, nil
}

func (r usecase) UpdateRealm(realmId uuid.UUID, realm dtos.UpdateRealmDTO) (*uuid.UUID, error) {
	r.logger.Info(fmt.Sprintf("Updating realm::: %s", realmId))

	updatedRealm := Realm{
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

	if err := r.db.Where("id = ?", realmId).Updates(&updatedRealm).Error; err != nil {
		return nil, err
	}

	return &realmId, nil
}

func (r usecase) DeleteRealm(realmId uuid.UUID) error {
	status := false
	deletedRealm := Realm{Enabled: &status}

	query := r.db.Where("id = ?", realmId).Updates(&deletedRealm)
	r.logger.Info(fmt.Sprintf("rows affected: %d", query.RowsAffected))

	if err := query.Error; err != nil {
		return err
	}

	return nil
}
