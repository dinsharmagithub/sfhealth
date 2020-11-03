package database

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dinsharmagithub/sfhealth/config"
)

var testCtx = context.Background()

func TestMockDb(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error creating sqlmock %v\n", err)
	}
	SetDbConnection(db)
	defer func() {
		SetDbConnection(nil)
	}()
	if GetDbConn() != db {
		db.Close()
		t.Fatalf("Database connection is not the one which is set")
	}

	// Check the connectivity.
	err = GetDbConn().Ping()
	if err != nil {
		t.Fatalf("DB not connectable: '%v'", err)
	}
}

func TestInitWithEmptyDb(t *testing.T) {
	cfg := config.Config{}
	err := Initialize(testCtx, cfg)
	defer SetDbConnection(nil)
	if err == nil {
		t.Fatalf("Error should have occured in this scenario\n")
	}
}

//TODO add more DB connection related tests
