// Package db provides functions to connect to the database and perform operations on it.
package db

import (
	"database/sql"
)

func Connect() error {
	db, err := sql.Open("sqlite3", "/app/data.db")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return Migrate(db)
}
