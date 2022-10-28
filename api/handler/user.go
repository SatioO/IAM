package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/internal/user"
	"github.com/satioO/iam/pkg/dtos"
	"github.com/satioO/iam/util"
)

type handler struct {
	usecase user.UserUsecase
}

func NewUserHandler(usecase user.UserUsecase) *handler {
	return &handler{usecase: usecase}
}

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userDetails, err := h.usecase.GetUsers(r.Context())

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, userDetails)
}

func (h handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	userDetails, err := h.usecase.GetUserDetails(r.Context())

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, userDetails)
}

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dtos.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.usecase.CreateUser(r.Context(), &user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, nil)
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user dtos.UpdateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(r)
	userId := uuid.MustParse(params["userId"])

	if err := h.usecase.UpdateUser(r.Context(), userId, &user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, nil)
}
