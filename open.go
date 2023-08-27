package sifdb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//Initiates a new table in the database. The table should be a struct with json tags
//to be able to easily unmarshal the content later.
//Example struct to give as table:
/*
type User struct {
	id        int     `json:"id"` -- mandatory field to have in every struct
	username  string  `json:"username"`
	password  string  `json:"password"`
}
*/
func Open(databasefile string) *sql.DB {
	// Open a database connection
	db, err := sql.Open("sqlite3", databasefile)
	if err != nil {
		panic(err)
	}

	return db
	// Now you can use the 'db' connection to perform database operations
}
