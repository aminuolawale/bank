package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/aminuolawale/bank/util"
	_ "github.com/lib/pq"
)


var testQueries *Queries
var	dbDriver = os.Getenv("POSTGRES_DB_DRIVER")
var	dbSource = os.Getenv("POSTGRES_DB_SOURCE")



func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db!", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
