package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var local *time.Location

func init() {
	local = time.FixedZone("CST", 8*3600) // 东八
	// local, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = local
}

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
