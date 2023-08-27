package sifdb

type encodedStruct struct {
	table_name   string
	field_names  []string
	field_values []string
	field_types  []string
}
