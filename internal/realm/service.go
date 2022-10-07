package realm

import (
	"github.com/google/uuid"
	"github.com/satioO/iam/pkg/dtos"
	"gorm.io/gorm"
)

type RealmUsecase interface {
	GetRealms() []dtos.ListRealmDTO
	CreateRealm(realm dtos.CreateRealmDTO) (*uuid.UUID, error)
}

type usecase struct {
	db *gorm.DB
}

func NewRealmUsecase(db *gorm.DB) *usecase {
	return &usecase{
		db,
	}
}

func (r usecase) GetRealms() []dtos.ListRealmDTO {
	var foundRealms []Realm
	r.db.Preload("Attributes").Find(&foundRealms)

	var result []dtos.ListRealmDTO
	for _, realm := range foundRealms {
		attributes := make(map[string]string)
		for _, attribute := range realm.Attributes {
			attributes[attribute.Name] = attribute.Value
		}

		result = append(result, dtos.ListRealmDTO{
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

func (r usecase) CreateRealm(body dtos.CreateRealmDTO) (*uuid.UUID, error) {
	defaultAttributes := map[string]string{
		"access_token_lifespan":  ACCESS_TOKEN_LIFESPAN,
		"refresh_token_lifespan": REFRESH_TOKEN_LIFESPAN,
	}

	for key, value := range body.Attributes {
		defaultAttributes[key] = value
	}

	var attributes []RealmAttribute
	for key, value := range defaultAttributes {
		attributes = append(attributes, RealmAttribute{
			Name:  key,
			Value: value,
		})
	}

	realm := &Realm{
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

	if err := r.db.Create(&realm).Error; err != nil {
		return nil, err
	}

	return &realm.ID, nil
}
