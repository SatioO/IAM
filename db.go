package main

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/satioO/iam/internal/client"
	"github.com/satioO/iam/internal/realm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN() string {
	return fmt.Sprintf("host=localhost user=admin password=admin dbname=iam port=5432 sslmode=disable")
}

func CreateDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Err(err).Msg("Error occurred while connecting with the database")
		return nil, err
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func AutoMigrateDB(db *gorm.DB) error {
	// Auto migrate database
	// Add new models here
	db.AutoMigrate(&realm.RealmAttribute{})
	db.AutoMigrate(&client.Client{})
	db.AutoMigrate(&realm.Realm{})
	return nil
}

func CloseDBConnection(conn *gorm.DB) {
	db, err := conn.DB()

	if err != nil {
		log.Err(err).Msg("Error occurred while closing a DB connection")
	}

	defer db.Close()
}
