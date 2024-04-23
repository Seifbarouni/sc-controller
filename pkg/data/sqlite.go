package data

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbName          = "processes.db"
	createProcessDb = `
  CREATE TABLE process (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    status VARCHAR(255) DEFAULT 'RUNNING',
    autorestart BOOLEAN DEFAULT TRUE,
    autostart BOOLEAN DEFAULT TRUE,
    working_directory VARCHAR(255) DEFAULT '.',
    stderr_logfile VARCHAR(255),
    stdout_logfile VARCHAR(255)
);
`
)

func Init() (*sql.DB, error) {
	if _, err := os.Stat("./" + dbName); errors.Is(err, os.ErrNotExist) {
		db, err := sql.Open("sqlite3", dbName)
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(createProcessDb)
		if err != nil {
			return nil, err
		}
	}
	return sql.Open("sqlite3", dbName)
}
