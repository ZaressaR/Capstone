package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///patient_profile?sslmode=disable"
	dbName   = "patient_profile"
)

var ctx = context.Background()
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	testdb, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testdb)
	os.Exit(m.Run())

}
