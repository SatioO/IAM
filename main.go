package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/satioO/iam/api"
	_ "github.com/satioO/iam/docs"
	"github.com/satioO/iam/pkg/log"
	"github.com/satioO/iam/pkg/trace"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gorm.io/gorm"
)

type server struct {
	router *mux.Router
	logger log.Factory
}

func (s *server) InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	s.logger = log.NewLogger(logger)
}

func (s *server) InitializeDatabase() *gorm.DB {
	// Initialize DB connection
	db, err := CreateDatabaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Migrate DB
	err = AutoMigrateDB(db)
	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}

func (s *server) InitializeRouter(db *gorm.DB) {
	// Initialize router
	router := api.NewMux(db, s.logger)
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

	srv.ListenAndServe()
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
	s.InitializeLogger()

	// Bootstrap tracer.
	tracer, err := trace.NewProvider(trace.ProviderConfig{
		JaegerEndpoint: "http://localhost:14268/api/traces",
		ServiceName:    "IAM",
		ServiceVersion: "1.0.0",
		Environment:    "dev",
		Disabled:       false,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	defer tracer.Close(ctx)

	s.InitializeRouter(db)
	s.Run("3000")

}
