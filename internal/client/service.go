package client

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/satioO/iam/internal/entities"
	"github.com/satioO/iam/pkg/dtos"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClientUsecase interface {
	GetClients() ([]dtos.ListClientsDTO, error)
	GetClientByID(clientId uuid.UUID) (*dtos.GetClientDTO, error)
	CreateClient(body dtos.CreateClientDTO) (*uuid.UUID, error)
	UpdateClient(clientId uuid.UUID, body dtos.UpdateClientDTO) (bool, error)
	DeleteClient(clientId uuid.UUID) (bool, error)
}

type usecase struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewClientUsecase(db *gorm.DB, logger *zap.Logger) *usecase {
	return &usecase{db, logger}
}

func (u usecase) GetClients() ([]dtos.ListClientsDTO, error) {
	u.logger.Info("Getting Clients:")

	var clients []entities.Client

	if err := u.db.Preload("Realm").Find(&clients).Error; err != nil {
		return nil, err
	}

	result := []dtos.ListClientsDTO{}

	for _, client := range clients {
		result = append(result, dtos.ListClientsDTO{
			ID:                       client.ID,
			ClientID:                 client.ClientID,
			Name:                     client.Name,
			Description:              client.Description,
			Protocol:                 client.Protocol,
			PublicClient:             client.PublicClient,
			StandardFlowEnabled:      client.StandardFlowEnabled,
			ImplicitFlowEnabled:      client.ImplicitFlowEnabled,
			DirectAccessGrantEnabled: client.DirectAccessGrantEnabled,
			ServiceAccountsEnabled:   client.ServiceAccountsEnabled,
			RootURL:                  client.RootURL,
			BaseURL:                  client.BaseURL,
			RedirectURIs:             client.RedirectURIs,
			Enabled:                  client.Enabled,
		})
	}

	return result, nil
}

func (u usecase) GetClientByID(clientId uuid.UUID) (*dtos.GetClientDTO, error) {
	u.logger.Info(fmt.Sprintf("Getting Client Details: %s", clientId))

	var foundClient entities.Client

	if err := u.db.Where("client_id = ?", clientId).First(&foundClient).Error; err != nil {
		return nil, err
	}

	return &dtos.GetClientDTO{
		ID:                       foundClient.ID,
		ClientID:                 foundClient.ClientID,
		Name:                     foundClient.Name,
		Description:              foundClient.Description,
		Protocol:                 foundClient.Protocol,
		PublicClient:             foundClient.PublicClient,
		StandardFlowEnabled:      foundClient.StandardFlowEnabled,
		ImplicitFlowEnabled:      foundClient.ImplicitFlowEnabled,
		DirectAccessGrantEnabled: foundClient.DirectAccessGrantEnabled,
		ServiceAccountsEnabled:   foundClient.ServiceAccountsEnabled,
		RootURL:                  foundClient.RootURL,
		BaseURL:                  foundClient.BaseURL,
		RedirectURIs:             foundClient.RedirectURIs,
		Enabled:                  foundClient.Enabled,
	}, nil
}

func (u usecase) CreateClient(body dtos.CreateClientDTO) (*uuid.UUID, error) {
	u.logger.Info(fmt.Sprintf("Creating Client: %s", body.ClientID))

	createdClient := &entities.Client{
		Name:                     body.Name,
		Description:              body.Description,
		Protocol:                 body.Protocol,
		ClientID:                 body.ClientID,
		RealmID:                  body.RealmID,
		PublicClient:             body.PublicClient,
		StandardFlowEnabled:      body.StandardFlowEnabled,
		ImplicitFlowEnabled:      body.ImplicitFlowEnabled,
		DirectAccessGrantEnabled: body.DirectAccessGrantEnabled,
		ServiceAccountsEnabled:   body.ServiceAccountsEnabled,
		RootURL:                  body.RootURL,
		BaseURL:                  body.BaseURL,
		RedirectURIs:             body.RedirectURIs,
		Enabled:                  body.Enabled,
	}

	if err := u.db.Create(&createdClient).Error; err != nil {
		return nil, err
	}

	return &createdClient.ID, nil
}

func (u usecase) UpdateClient(clientId uuid.UUID, body dtos.UpdateClientDTO) (bool, error) {
	u.logger.Info(fmt.Sprintf("Updating Client: %s", clientId))

	updatedClient := &entities.Client{
		Name:                     body.Name,
		Description:              body.Description,
		Protocol:                 body.Protocol,
		PublicClient:             body.PublicClient,
		StandardFlowEnabled:      body.StandardFlowEnabled,
		ImplicitFlowEnabled:      body.ImplicitFlowEnabled,
		DirectAccessGrantEnabled: body.DirectAccessGrantEnabled,
		ServiceAccountsEnabled:   body.ServiceAccountsEnabled,
		RootURL:                  body.RootURL,
		BaseURL:                  body.BaseURL,
		RedirectURIs:             body.RedirectURIs,
	}

	if err := u.db.Where("client_id = ?", clientId).Updates(&updatedClient).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (u usecase) DeleteClient(clientId uuid.UUID) (bool, error) {
	u.logger.Info(fmt.Sprintf("Deleting Client: %s", clientId))
	if err := u.db.Where("client_id = ?", clientId).Update("enabled", false).Error; err != nil {
		return false, err
	}

	return true, nil
}
