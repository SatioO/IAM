package realm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ACCESS_TOKEN_LIFESPAN  = "300"
	REFRESH_TOKEN_LIFESPAN = "600"
)

type Realm struct {
	ID                      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name                    string    `gorm:"unique;not null"`
	DisplayName             string
	Logo                    string
	SupportURL              string
	SupportEmail            string
	DuplicateEmailAllowed   *bool     `gorm:"default:false"`
	DuplicatePhoneAllowed   *bool     `gorm:"default:false"`
	RegisterEmailAsUsername *bool     `gorm:"default:true"`
	RegisterPhoneAsUsername *bool     `gorm:"default:false"`
	Enabled                 *bool     `gorm:"not null"`
	CreatedAt               time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt               time.Time `gorm:"autoUpdateTime:milli"`
	Attributes              []RealmAttribute
}

// TableName overrides the table name used by User to `profiles`
func (Realm) TableName() string {
	return "REALM"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *Realm) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type RealmAttribute struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"not null"`
	RealmID   string
	Value     string
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

// TableName overrides the table name used by User to `profiles`
func (RealmAttribute) TableName() string {
	return "REALM_ATTRIBUTE"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *RealmAttribute) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
