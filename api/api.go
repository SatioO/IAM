package api

import (
	"net/http"

	chimid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/api/handler"
	"github.com/satioO/iam/internal/realm"
	"gorm.io/gorm"
)

func NewMux(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	r.Use(chimid.RequestID)
	r.Use(chimid.RealIP)
	r.Use(chimid.Recoverer)

	// Realm Management
	realmUsecase := realm.NewRealmUsecase(db)
	realmHandler := handler.NewRealmHandler(realmUsecase)

	r.HandleFunc("/realm", realmHandler.GetRealms).Methods(http.MethodGet)
	r.HandleFunc("/realm", realmHandler.CreateRealm).Methods(http.MethodPost)

	return r
}
