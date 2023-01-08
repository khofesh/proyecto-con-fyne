package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepository(db)
	os.Exit(m.Run())
}
