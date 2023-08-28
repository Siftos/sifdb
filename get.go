package sifdb

//small test
import (
	"database/sql"
	"fmt"
)

// sifdb.Get is what you use to fetch from the database. It takes
// a struct containing what fields you want to filter on. Example program:
/*
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db := sifdb.Open("test.sqlite")
	if err := sifdb.Create(User{}, db); err != nil {
		fmt.Println(err)
		return
	}
	user := User{}
	user.Id = 2
	if user, err := getUser(user, db); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(user)
	}
}

func getUser(user User, db *sql.DB) (User, error) {
	jsonString, err := sifdb.Get(user, db)
	if err != nil {
		return User{}, err
	}
	if len(jsonString) == 0 {
		return User{}, errors.New("no user found")
	}
	if len(jsonString) > 1 {
		return User{}, errors.New("found too many users")
	}
	json.Unmarshal([]byte(jsonString[0]), &user)
	return user, nil
}
*/
func Get(database_struct interface{}, db *sql.DB) ([]string, error) {
	encoded_struct, err := encodeToArray(database_struct)
	if err != nil {
		return []string{}, err
	}

	sqlStatement := createSelectSQLStatement(encoded_struct)
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

func createSelectSQLStatement(encoded_struct encodedStruct) string {
	sqlStatement := "SELECT * FROM " + encoded_struct.table_name
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

func assignTempEncodedStruct(vals []interface{}, encoded_struct *encodedStruct) error {
	for i, v := range vals {
		tmp_string := fmt.Sprintf("%s", v)
		encoded_struct.field_values[i] = tmp_string[1:]
	}
	return nil
}
