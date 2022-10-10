package handler

import (
	"github.com/satioO/iam/internal/client"
	"go.uber.org/zap"
)

type clientHandler struct {
	usecase client.ClientUsecase
	logger  *zap.Logger
}

func NewClientHandler(usecase client.ClientUsecase, logger *zap.Logger) *clientHandler {
	return &clientHandler{usecase, logger}
}
