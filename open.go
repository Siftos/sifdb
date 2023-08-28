package sifdb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Opens a new database, which points to an sqlite3 database file.

func Open(databasefile string) *sql.DB {
	// Open a database connection
	db, err := sql.Open("sqlite3", databasefile)
	if err != nil {
		panic(err)
	}

	return db
	// Now you can use the 'db' connection to perform database operations
}
