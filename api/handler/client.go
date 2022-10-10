package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/internal/client"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/util"
	"go.uber.org/zap"
)

type clientHandler struct {
	usecase client.ClientUsecase
	logger  *zap.Logger
}

func NewClientHandler(usecase client.ClientUsecase, logger *zap.Logger) *clientHandler {
	return &clientHandler{usecase, logger}
}

func (h clientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	foundClient, err := h.usecase.GetClients()

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundClient)
}

func (h clientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	foundClient, err := h.usecase.GetClientByID(uuid.MustParse(params["clientId"]))

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundClient)
}

func (h clientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateClientDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdClient, err := h.usecase.CreateClient(body)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, createdClient)
}

func (h clientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var body dtos.UpdateClientDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(r)
	updatedClient, err := h.usecase.UpdateClient(uuid.MustParse(params["clientId"]), body)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, updatedClient)
}

func (h clientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deletedClient, err := h.usecase.DeleteClient(uuid.MustParse(params["clientId"]))

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, deletedClient)
}
