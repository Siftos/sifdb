package sifdb

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func encodeToArray(database_struct interface{}) (encodedStruct, error) {
	value := reflect.ValueOf(database_struct)
	table_name := fmt.Sprintf("%T", database_struct)
	name_array := []string{}
	value_array := []string{}
	type_array := []string{}

	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			fieldName := value.Type().Field(i).Name
			fieldValue := value.Field(i).Interface()
			name_array = append(name_array, fieldName)
			if tmp, ok := fieldValue.(int); ok {
				value_array = append(value_array, strconv.Itoa(tmp))
				type_array = append(type_array, "int")
				continue
			}
			value_array = append(value_array, fmt.Sprintf("%s", fieldValue))
			type_array = append(type_array, fmt.Sprintf("%T", fieldValue))
		}
		return encodedStruct{table_name, name_array, value_array, type_array}, nil
	}
	return encodedStruct{}, errors.New("the given database_struct is not of the type struct")
}
