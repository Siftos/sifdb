package sifdb

import (
	"database/sql"
	"fmt"
)

func GetCustom(database_struct interface{}, db *sql.DB, where_statement string) ([]string, error) {
	encoded_struct, err := encodeToArray(database_struct)
	if err != nil {
		return []string{}, err
	}

	sqlStatement := createSelectSQLStatementWithCustomWhere(encoded_struct, where_statement)
	fmt.Println(sqlStatement)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []string{}, err
	}

	cols, err := rows.Columns() // Remember to check err afterwards
	if err != nil {
		return []string{}, err
	}

	vals := make([]interface{}, len(cols))
	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}

	result := []string{}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			return []string{}, err
		}
		temp_encoded_struct, err := encodeToArray(database_struct)
		if err != nil {
			return []string{}, err
		}
		if err := assignTempEncodedStruct(vals, &temp_encoded_struct); err != nil {
			return []string{}, err
		}
		result = append(result, createJson(temp_encoded_struct))
	}

	return result, nil
}

func createSelectSQLStatementWithCustomWhere(encoded_struct encodedStruct, custom_where string) string {
	sqlStatement := "SELECT * FROM " + encoded_struct.table_name
	return sqlStatement + " WHERE " + custom_where
}
