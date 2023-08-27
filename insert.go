package sifdb

import (
	"database/sql"
	"fmt"
)

func Insert(database_struct interface{}, db *sql.DB) error {
	encoded_struct, err := encodeToArray(database_struct)
	if err != nil {
		return err
	}
	field_names := "("
	field_values := "("
	for i := 1; i < len(encoded_struct.field_names); i++ {
		field_names += encoded_struct.field_names[i]
		field_values += "?"
		if i < len(encoded_struct.field_names)-1 {
			field_names += ", "
			field_values += ", "
		}
	}
	field_names += ")"
	field_values += ")"
	sqlStatement := "INSERT INTO " + encoded_struct.table_name + " " + field_names + " VALUES " + field_values

	new_field_values := []interface{}{}
	for i := 1; i < len(encoded_struct.field_values); i++ {
		new_field_values = append(new_field_values, encoded_struct.field_values[i])
	}

	fmt.Println(sqlStatement)
	_, err = db.Exec(sqlStatement, new_field_values...)

	if err != nil {
		return err
	}

	return nil
}
