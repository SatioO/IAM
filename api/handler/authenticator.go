package handler

import (
	"net/http"

	"github.com/satioO/iam/util"
)

type authenticator struct{}

func NewAuthenticator() *authenticator {
	return &authenticator{}
}

func (a authenticator) GetAuthenticators(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (a authenticator) GetAuthenticatorDetails(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (a authenticator) CreateAuthenticator(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (a authenticator) UpdateAuthenticator(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}

func (a authenticator) DeleteAuthenticator(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusNotImplemented, nil)
}
