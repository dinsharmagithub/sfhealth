package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/dinsharmagithub/sfhealth/config"
)

var dbConnection *sql.DB

// Initialize : iniatialize db connection
func Initialize(ctx context.Context, cfg config.Config) error {

	// cfg.ConnStr = "user=dineshs dbname=dineshs password=ChangeIt host=127.0.0.1 sslmode=disable"
	// cfg.Driver = "postgres"
	dbConn, err := sql.Open(cfg.Driver, cfg.ConnStr)
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	//TODO: Remove hardcoding below and get value from config file
	dbConn.SetMaxIdleConns(3)
	dbConn.SetMaxOpenConns(3)

	log.Printf("DB connection successful")

	dbConnection = dbConn

	return nil
}

// CloseDbConn closes the database handle.
func CloseDbConn(ctx context.Context) error {
	err := dbConnection.Close()
	if err != nil {
		log.Fatalf("Failed to close the db connection: %v", err)
		return err
	}

	return nil
}

// GetDbConn returns global *sql.DB handle.
func GetDbConn() *sql.DB {
	if dbConnection == nil {
		errorMsg := "GetDbConn called before initializing the database connection"
		panic(errorMsg)
	}

	return dbConnection
}

//SetDbConnection sets a db connection
func SetDbConnection(db *sql.DB) {
	dbConnection = db
}
