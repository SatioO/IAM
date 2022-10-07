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
	r.db.Find(&foundRealms)

	var result []dtos.ListRealmDTO
	for _, realm := range foundRealms {
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
		})
	}

	return result
}

func (r usecase) CreateRealm(body dtos.CreateRealmDTO) (*uuid.UUID, error) {
	var attributes []RealmAttribute

	for key, value := range body.Attributes {
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

	createdRealm := r.db.Create(&realm)

	if createdRealm.Error != nil {
		return nil, createdRealm.Error
	}

	return &realm.ID, nil
}
