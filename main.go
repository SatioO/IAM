package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/satioO/iam/api"
)

type server struct {
	router *mux.Router
}

func (s *server) Initialize() {
	// Initialize DB connection
	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Err(err).Msg("Error connecting DB")
	}

	// Migrate DB
	err = AutoMigrateDB(db)
	if err != nil {
		log.Err(err).Msg("Error occurred while connecting with the database")
	}

	// Initialize router
	s.router = api.NewMux(db)
}

func (s *server) Run(servicePort string) {
	srv := &http.Server{
		Addr:    ":" + servicePort,
		Handler: s.router,
	}
	fmt.Printf("App started listening on PORT %s\n", servicePort)
	log.Err(srv.ListenAndServe()).Msg("Error running service")
}

func main() {
	s := server{}
	s.Initialize()
	s.Run("3000")
}
