package realm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Realm struct {
	ID                      uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name                    string    `gorm:"unique;not null" json:"name" binding:"required"`
	DisplayName             string    `json:"display_name" binding:"required"`
	DuplicateEmailAllowed   bool      `gorm:"default:false" json:"duplicate_email_allowed"`
	DuplicatePhoneAllowed   bool      `gorm:"default:false" json:"duplicate_phone_allowed"`
	RegisterEmailAsUsername bool      `gorm:"default:true" json:"register_email_as_username"`
	RegisterPhoneAsUsername bool      `gorm:"default:false" json:"register_phone_as_username"`
	Enabled                 bool      `gorm:"not null" json:"enabled"`
	CreatedAt               time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt               time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *Realm) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
