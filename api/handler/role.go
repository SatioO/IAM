package handler

import (
	"net/http"

	"github.com/satioO/iam/util"
)

type role struct{}

func NewRoleHandler() *role {
	return &role{}
}

func (u role) GetRoles(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (u role) GetRoleDetails(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (u role) CreateRole(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (u role) UpdateRole(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (u role) DeleteRole(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}
