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

// List clients godoc
// @Summary      List of clients
// @Description  This operation is to get list of clients
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {array}   dtos.ListClientsDTO
// @Router       /client [get]
func (h clientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	foundClient, err := h.usecase.GetClients()

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundClient)
}

// Get client details godoc
// @Summary      Get details of client
// @Description  This operation is to get details of the client
// @Param        clientId    path     string  false  "client identifier"
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {object}   dtos.GetClientDTO
// @Router       /client/{clientId} [get]
func (h clientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	foundClient, err := h.usecase.GetClientByID(uuid.MustParse(params["clientId"]))

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundClient)
}

// Creates Client godoc
// @Summary      Creates Client
// @Description  This operation is to create new client
// @Param client body dtos.CreateClientDTO true "client data"
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {string}   uuid.UUID
// @Router       /client [post]
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

// Update Client Details godoc
// @Summary      Update client
// @Description  This operation is to update new client
// @Param client body dtos.UpdateClientDTO true "client data"
// @Param clientId query string true "client identifier"
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {string}   uuid.UUID
// @Router       /client/{clientId} [put]
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

// Delete Client godoc
// @Summary      Delete Client
// @Description  This operation is to delete client
// @Param clientId query string true "client identifier"
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      200  {boolean}   bool
// @Router       /client/{clientId} [delete]
func (h clientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deletedClient, err := h.usecase.DeleteClient(uuid.MustParse(params["clientId"]))

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, deletedClient)
}
