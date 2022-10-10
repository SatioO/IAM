package client

import (
	"errors"

	"github.com/google/uuid"
	"github.com/satioO/iam/pkg/dtos"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type usecase struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewClientUsecase(db *gorm.DB, logger *zap.Logger) *usecase {
	return &usecase{db, logger}
}

func (u usecase) GetClients() (*dtos.ListClientsDTO, error) {
	return nil, errors.New("Not Implmeneted")
}

func (u usecase) GetClientByID(clientId uuid.UUID) (*dtos.GetClientDTO, error) {
	return nil, errors.New("Not Implmeneted")
}

func (u usecase) CreateClient(body dtos.CreateClientDTO) (*uuid.UUID, error) {
	return nil, errors.New("Not Implmeneted")
}

func (u usecase) UpdateClient(clientId uuid.UUID, body dtos.UpdateClientDTO) (*uuid.UUID, error) {
	return nil, errors.New("Not Implmeneted")
}

func (u usecase) DeleteClient(clientId uuid.UUID) (*bool, error) {
	return nil, errors.New("Not Implmeneted")
}
