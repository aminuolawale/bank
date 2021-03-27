package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/aminuolawale/bank/util"
	_ "github.com/lib/pq"
)
var	dbDriver = os.Getenv("POSTGRES_DB_DRIVER")
var	dbSource = os.Getenv("POSTGRES_DB_SOURCE")


var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db!", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
