package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/satioO/iam/api"
	_ "github.com/satioO/iam/docs"
	"github.com/satioO/iam/pkg/trace"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
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

func (s *server) InitializeTracer() {

}

func (s *server) InitializeDatabase() *gorm.DB {
	// Initialize DB connection
	db, err := CreateDatabaseConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
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
	router := api.NewMux(db, logger)
	router.Use(otelmux.Middleware("IAM"))

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	s.router = router
}

func (s server) Run(servicePort string) {
	srv := &http.Server{
		Addr:    ":" + servicePort,
		Handler: s.router,
	}
	s.logger.Info(fmt.Sprintf("App started listening on PORT %s", servicePort))
	s.logger.Fatal(srv.ListenAndServe().Error())
}

// @title Identity Management Server
// @version 1.0
// @description These apis are built to handle identity management.

// @contact.name
// @contact.email vaibhav.satam29991@gmail.com
func main() {
	ctx := context.Background()

	s := server{}
	db := s.InitializeDatabase()
	logger := s.InitializeLogger()

	// Bootstrap tracer.
	tracer, err := trace.NewProvider(trace.ProviderConfig{
		JaegerEndpoint: "http://localhost:14268/api/traces",
		ServiceName:    "IAM",
		ServiceVersion: "1.0.0",
		Environment:    "dev",
		Disabled:       false,
	})

	if err != nil {
		logger.Fatal(err.Error())
	}

	defer tracer.Close(ctx)

	s.InitializeTracer()
	s.InitializeRouter(db, logger)
	s.Run("3000")

}
