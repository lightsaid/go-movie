package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://postgres:abc123@localhost:5432/db_movie?sslmode=disable")
	if err != nil {
		log.Fatalf("init db error %q", err.Error())
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
