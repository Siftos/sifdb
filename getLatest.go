package sifdb

import (
	"database/sql"
	"fmt"
	"strconv"
)

func GetLatest(database_struct interface{}, amount int, db *sql.DB) ([]string, error) {
	encoded_struct, err := encodeToArray(database_struct)
	if err != nil {
		return []string{}, err
	}

	sqlStatement := createSelectSQLStatement(encoded_struct)
	sqlStatement = addOrderBy(sqlStatement, amount)

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

func addOrderBy(sqlStatement string, amount int) string {
	return sqlStatement + " ORDER BY id DESC LIMIT " + strconv.Itoa(amount)
}
