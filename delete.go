package sifdb

import "database/sql"

func Delete(database_struct interface{}, db *sql.DB) error {
	encoded_struct, err := encodeToArray(database_struct)
	if err != nil {
		return err
	}

	sqlStatement := createDeleteSQLStatement(encoded_struct)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

func createDeleteSQLStatement(encoded_struct encodedStruct) string {
	sqlStatement := "DELETE FROM " + encoded_struct.table_name
	where_is_used := false
	where_content := ""

	for i := 0; i < len(encoded_struct.field_names); i++ {
		if encoded_struct.field_values[i] == "" {
			continue
		}

		if encoded_struct.field_values[i] == "0" && i == 0 {
			continue
		}

		where_is_used = true
		if where_content != "" {
			where_content += " AND "
		}

		if encoded_struct.field_types[i] == "string" {
			where_content += encoded_struct.field_names[i] + " = \"" + encoded_struct.field_values[i] + "\""
			continue
		}

		where_content += encoded_struct.field_names[i] + " = " + encoded_struct.field_values[i]
	}

	if where_is_used {
		sqlStatement += " WHERE " + where_content
	}
	return sqlStatement
}
