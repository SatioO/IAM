package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/satioO/iam/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type server struct {
	router *mux.Router
	logger *zap.Logger
}

func (s *server) InitializeLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	s.logger = logger
	return logger
}

func (s *server) InitializeDatabase() *gorm.DB {
	// Initialize DB connection
	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Err(err).Msg("Error connecting DB")
	}

	// Migrate DB
	err = AutoMigrateDB(db)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	return db
}

func (s *server) InitializeRouter(db *gorm.DB, logger *zap.Logger) {
	// Initialize router
	s.router = api.NewMux(db, logger)
}

func (s server) Run(servicePort string) {
	srv := &http.Server{
		Addr:    ":" + servicePort,
		Handler: s.router,
	}
	s.logger.Info(fmt.Sprintf("App started listening on PORT %s", servicePort))
	s.logger.Fatal(srv.ListenAndServe().Error())
}

func main() {
	s := server{}
	db := s.InitializeDatabase()
	logger := s.InitializeLogger()
	s.InitializeRouter(db, logger)
	s.Run("3000")
}
