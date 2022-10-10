package client

import (
	"github.com/google/uuid"
	"github.com/satioO/iam/pkg/dtos"
	"gorm.io/gorm"
)

type ClientUsecase interface {
	GetClients() (*dtos.GetClientsDTO, error)
	GetClientByID(clientId uuid.UUID) (*dtos.GetClientDTO, error)
	CreateClient(body dtos.CreateClientDTO) (*uuid.UUID, error)
	UpdateClient(clientId uuid.UUID, body dtos.UpdateClientDTO) (*uuid.UUID, error)
	DeleteClient(clientId uuid.UUID) (*bool, error)
}

type Client struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key"`
	ClientID                 string    `gorm:"not null"`
	RealmID                  uuid.UUID
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
