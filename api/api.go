package api

import (
	"net/http"

	chimid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/api/handler"
	"github.com/satioO/iam/internal/client"
	"github.com/satioO/iam/internal/realm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewMux(db *gorm.DB, logger *zap.Logger) *mux.Router {
	r := mux.NewRouter()

	r.Use(chimid.RequestID)
	r.Use(chimid.RealIP)
	r.Use(chimid.Recoverer)

	// Realm Management
	realmUsecase := realm.NewRealmUsecase(db, logger)
	realmHandler := handler.NewRealmHandler(realmUsecase, logger)

	r.HandleFunc("/realm", realmHandler.GetRealms).Methods(http.MethodGet)
	r.HandleFunc("/realm/{realmId}", realmHandler.GetRealmByID).Methods(http.MethodGet)
	r.HandleFunc("/realm", realmHandler.CreateRealm).Methods(http.MethodPost)
	r.HandleFunc("/realm/{realmId}", realmHandler.UpdateRealm).Methods(http.MethodPut)
	r.HandleFunc("/realm/{realmId}", realmHandler.DeleteRealm).Methods(http.MethodDelete)

	// Client Management
	clientUsecase := client.NewClientUsecase(db, logger)
	clientHandler := handler.NewClientHandler(clientUsecase, logger)

	r.HandleFunc("/client", clientHandler.GetClients).Methods(http.MethodGet)
	r.HandleFunc("/client/{clientId}", clientHandler.GetClientByID).Methods(http.MethodGet)
	r.HandleFunc("/client", clientHandler.CreateClient).Methods(http.MethodPost)
	r.HandleFunc("/client/{clientId}", clientHandler.UpdateClient).Methods(http.MethodPut)
	r.HandleFunc("/client/{clientId}", clientHandler.DeleteClient).Methods(http.MethodDelete)

	return r
}
