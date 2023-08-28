package sifdb

import (
	"database/sql"
	"errors"
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
func Create(table interface{}, db *sql.DB) error {
	encoded_struct, err := encodeToArray(table)
	standard_err_msg := "error while trying to create table " + encoded_struct.table_name + ": "
	if err != nil {
		return err
	}

	id := encoded_struct.field_names[0]
	if id != "id" && id != "Id" && id != "ID" {
		return errors.New(standard_err_msg + "the first field needs to be named id, it is currently named " + id)
	}

	id_type := encoded_struct.field_types[0]
	if id_type != "int" {
		return errors.New(standard_err_msg + "the id field must be of the type 'int' it is currently ")
	}

	createTableStatement := `CREATE TABLE IF NOT EXISTS ` + encoded_struct.table_name + `(
		` + id + ` INTEGER PRIMARY KEY AUTOINCREMENT, `

	for i := 1; i < len(encoded_struct.field_names); i++ {
		if encoded_struct.field_types[i] != "int" && encoded_struct.field_types[i] != "string" {
			return errors.New(standard_err_msg + "type " + encoded_struct.field_types[i] + " is not supported")
		}

		createTableStatement += encoded_struct.field_names[i] + ` TEXT`
		if i < len(encoded_struct.field_names)-1 {
			createTableStatement += `, `
		}
	}
	createTableStatement += `);`
	if _, err := db.Exec(createTableStatement); err != nil {
		return err
	}
	return nil
}
