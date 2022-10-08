package handler

import (
	"encoding/json"
	"net/http"

	"github.com/satioO/iam/internal/realm"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/util"
	"go.uber.org/zap"
)

type handler struct {
	realm  realm.RealmUsecase
	logger *zap.Logger
}

func NewRealmHandler(usecase realm.RealmUsecase, logger *zap.Logger) *handler {
	return &handler{
		realm:  usecase,
		logger: logger,
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
	return
}

func (r handler) UpdateRealm(res http.ResponseWriter, req *http.Request) {}
