package client

import (
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
