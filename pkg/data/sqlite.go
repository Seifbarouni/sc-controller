package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbName = "processes.db"
)

func Init() (*sql.DB, error) {
	return sql.Open("sqlite3", dbName)
}
