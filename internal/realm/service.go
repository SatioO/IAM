package realm

import (
	"gorm.io/gorm"
)

type RealmUsecase interface {
	GetRealms() []Realm
	CreateRealm(realm Realm) (*Realm, error)
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

func (r usecase) CreateRealm(realm Realm) (*Realm, error) {
	result := r.db.Create(&realm)

	if result.Error != nil {
		return nil, result.Error
	}

	return &realm, nil
}
