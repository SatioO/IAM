package realm

import (
	"github.com/google/uuid"
	"github.com/satioO/iam/pkg/dtos"
	"gorm.io/gorm"
)

type RealmUsecase interface {
	GetRealms() []Realm
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

func (r usecase) GetRealms() []Realm {
	var realms []Realm

	r.db.Find(&realms)

	return realms
}

func (r usecase) CreateRealm(body dtos.CreateRealmDTO) (*uuid.UUID, error) {
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
	}

	createdRealm := r.db.Create(&realm)

	if createdRealm.Error != nil {
		return nil, createdRealm.Error
	}

	return &realm.ID, nil
}
