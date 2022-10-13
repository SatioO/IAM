package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key"`
	ClientID                 string    `gorm:"not null"`
	RealmID                  uuid.UUID `gorm:"column:realm_id"`
	Realm                    Realm
	Name                     string `gorm:"not null"`
	Description              string
	Protocol                 string `gorm:"default:openid-connect;not null"`
	PublicClient             *bool  `gorm:"default:false;not null"`
	StandardFlowEnabled      *bool  `gorm:"default:false;not null"`
	ImplicitFlowEnabled      *bool  `gorm:"default:false;not null"`
	DirectAccessGrantEnabled *bool  `gorm:"default:false;not null"`
	ServiceAccountsEnabled   *bool  `gorm:"default:false;not null"`
	RootURL                  string
	BaseURL                  string
	RedirectURIs             string
	Enabled                  *bool `gorm:"default:false;not null"`
}

func (Client) TableName() string {
	return "CLIENT"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *Client) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
