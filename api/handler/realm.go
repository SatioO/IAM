package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/internal/realm"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/pkg/log"
	"github.com/satioO/iam/util"
)

type realmHandler struct {
	realm  realm.RealmUsecase
	logger log.Factory
}

func NewRealmHandler(usecase realm.RealmUsecase, logger log.Factory) *realmHandler {
	return &realmHandler{
		realm:  usecase,
		logger: logger,
	}
}

// List Realms godoc
// @Summary      List of realms
// @Description  This operation is to get list of realms
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {array}   dtos.ListRealmDTO
// @Router       /realm [get]
func (h realmHandler) GetRealms(w http.ResponseWriter, r *http.Request) {
	foundRealms, err := h.realm.GetRealms(r.Context())

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundRealms)
}

// Get Realm Details godoc
// @Summary      Get details of realm
// @Description  This operation is to get details of the realm
// @Param        realmId    path     string  false  "realm identifier"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {object}   dtos.GetRealmDTO
// @Router       /realm/{realmId} [get]
func (h realmHandler) GetRealmByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	foundRealm, err := h.realm.GetRealmByID(uuid.MustParse(params["realmId"]))

	if err != nil {
		util.RespondWithError(w, http.StatusNotFound, "No realm found")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, foundRealm)
}

// Create Realm Details godoc
// @Summary      Create Realm
// @Description  This operation is to create new realm
// @Param realm body dtos.CreateRealmDTO true "realm data"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {string}   uuid.UUID
// @Router       /realm [post]
func (h realmHandler) CreateRealm(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateRealmDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdRealm, err := h.realm.CreateRealm(body)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, createdRealm)
}

// Update Realm Details godoc
// @Summary      Update Realm
// @Description  This operation is to update new realm
// @Param realm body dtos.UpdateRealmDTO true "realm data"
// @Param realmId query string true "realm identifier"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {string}   uuid.UUID
// @Router       /realm/{realmId} [put]
func (h realmHandler) UpdateRealm(w http.ResponseWriter, r *http.Request) {
	var body dtos.UpdateRealmDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(r)

	updatedRealm, err := h.realm.UpdateRealm(uuid.MustParse(params["realmId"]), body)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, updatedRealm)
}

// Delete Realm godoc
// @Summary      Delete Realm
// @Description  This operation is to delete realm
// @Param realmId query string true "realm identifier"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {boolean}   bool
// @Router       /realm/{realmId} [delete]
func (h realmHandler) DeleteRealm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if _, err := h.realm.DeleteRealm(uuid.MustParse(params["realmId"])); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, true)
}
