package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:Avintiv@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect ot db: ", err)

	}
	testQueries = New(testDb)
	os.Exit(m.Run())

	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	// connPool, err := pgxpool.New(context.Background(), config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	// testStore = NewStore(connPool)
	// os.Exit(m.Run())
}
