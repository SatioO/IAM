package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/internal/realm"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/util"
	"go.uber.org/zap"
)

type realmHandler struct {
	realm  realm.RealmUsecase
	logger *zap.Logger
}

func NewRealmHandler(usecase realm.RealmUsecase, logger *zap.Logger) *realmHandler {
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
func (r realmHandler) GetRealms(res http.ResponseWriter, req *http.Request) {
	foundRealms := r.realm.GetRealms()

	if len(foundRealms) == 0 {
		util.RespondWithError(res, http.StatusNotFound, "No realms found")
		return
	}

	util.RespondWithJSON(res, http.StatusOK, foundRealms)
}

// Get Realm Details godoc
// @Summary      Get details of realm
// @Description  This operation is to get details of the realm
// @Param        realmId    path     string  false  "realm identifier"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {array}   dtos.GetRealmDTO
// @Router       /realm/{realmId} [get]
func (r realmHandler) GetRealmByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	foundRealm, err := r.realm.GetRealmByID(uuid.MustParse(params["realmId"]))

	if err != nil {
		util.RespondWithError(res, http.StatusNotFound, "No realm found")
		return
	}

	util.RespondWithJSON(res, http.StatusOK, foundRealm)
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
func (r realmHandler) CreateRealm(res http.ResponseWriter, req *http.Request) {
	var body dtos.CreateRealmDTO

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	createdRealm, err := r.realm.CreateRealm(body)

	if err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(res, http.StatusOK, createdRealm)
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
func (r realmHandler) UpdateRealm(res http.ResponseWriter, req *http.Request) {
	var body dtos.UpdateRealmDTO

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(req)

	updatedRealm, err := r.realm.UpdateRealm(uuid.MustParse(params["realmId"]), body)

	if err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(res, http.StatusOK, updatedRealm)
}

// Delete Realm godoc
// @Summary      Delete Realm
// @Description  This operation is to delete new realm
// @Param realmId query string true "realm identifier"
// @Tags         Realm
// @Accept       json
// @Produce      json
// @Success      200  {boolean}   bool
// @Router       /realm/{realmId} [delete]
func (r realmHandler) DeleteRealm(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	if err := r.realm.DeleteRealm(uuid.MustParse(params["realmId"])); err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(res, http.StatusOK, true)
}
