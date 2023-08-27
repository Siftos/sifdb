package sifdb

func createJson(encoded_struct encodedStruct) string {
	resultJson := "{"
	for i := 0; i < len(encoded_struct.field_names); i++ {
		resultJson += "\"" + encoded_struct.field_names[i] + "\":"
		if encoded_struct.field_types[i] == "int" {
			resultJson += encoded_struct.field_values[i] + ", "
			continue
		}
		resultJson += "\"" + encoded_struct.field_values[i] + "\""
		if i < len(encoded_struct.field_names)-1 {
			resultJson += ", "
		}
	}
	resultJson += "}"
	return resultJson
}
