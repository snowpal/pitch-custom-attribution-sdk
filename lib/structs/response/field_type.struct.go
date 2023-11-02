package response

type FieldTypes struct {
	FieldTypes []FieldType `json:"fieldTypes"`
}

type FieldType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
