package handler

import (
	"encoding/json"
	"net/http"

	"github.com/satioO/iam/internal/realm"
	"github.com/satioO/iam/util"
)

type handler struct {
	realm realm.RealmUsecase
}

func NewRealmHandler(usecase realm.RealmUsecase) *handler {
	return &handler{
		realm: usecase,
	}
}

func (r handler) GetRealms(res http.ResponseWriter, req *http.Request) {
	foundRealms := r.realm.GetRealms()

	if len(foundRealms) == 0 {
		util.RespondWithError(res, http.StatusNotFound, "No realms found")
		return
	}

	util.RespondWithJSON(res, http.StatusOK, foundRealms)
	return
}

func (r handler) CreateRealm(res http.ResponseWriter, req *http.Request) {
	var body realm.Realm
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, "Failed to decode request body")
		return
	}

	createdRealm, err := r.realm.CreateRealm(body)

	if err != nil {
		util.RespondWithError(res, http.StatusInternalServerError, "Failed to create realm")
		return
	}

	util.RespondWithJSON(res, http.StatusOK, createdRealm)
	return
}
