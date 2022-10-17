package api

import (
	"net/http"

	chimid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/satioO/iam/api/handler"
	"github.com/satioO/iam/internal/client"
	"github.com/satioO/iam/internal/realm"
	"github.com/satioO/iam/pkg/trace"
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

	r.HandleFunc("/realm", trace.HTTPHandlerFunc(realmHandler.GetRealms, "GetRealms")).Methods(http.MethodGet)
	r.HandleFunc("/realm/{realmId}", trace.HTTPHandlerFunc(realmHandler.GetRealmByID, "GetRealmByID")).Methods(http.MethodGet)
	r.HandleFunc("/realm", trace.HTTPHandlerFunc(realmHandler.CreateRealm, "CreateRealm")).Methods(http.MethodPost)
	r.HandleFunc("/realm/{realmId}", trace.HTTPHandlerFunc(realmHandler.UpdateRealm, "UpdateRealm")).Methods(http.MethodPut)
	r.HandleFunc("/realm/{realmId}", trace.HTTPHandlerFunc(realmHandler.DeleteRealm, "DeleteRealm")).Methods(http.MethodDelete)

	// Client Management
	clientUsecase := client.NewClientUsecase(db, logger)
	clientHandler := handler.NewClientHandler(clientUsecase, logger)

	r.HandleFunc("/client", trace.HTTPHandlerFunc(clientHandler.GetClients, "GetClients")).Methods(http.MethodGet)
	r.HandleFunc("/client/{clientId}", trace.HTTPHandlerFunc(clientHandler.GetClientByID, "GetClientByID")).Methods(http.MethodGet)
	r.HandleFunc("/client", trace.HTTPHandlerFunc(clientHandler.CreateClient, "CreateClient")).Methods(http.MethodPost)
	r.HandleFunc("/client/{clientId}", trace.HTTPHandlerFunc(clientHandler.UpdateClient, "UpdateClient")).Methods(http.MethodPut)
	r.HandleFunc("/client/{clientId}", trace.HTTPHandlerFunc(clientHandler.DeleteClient, "DeleteClient")).Methods(http.MethodDelete)

	return r
}
