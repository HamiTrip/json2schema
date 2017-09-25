package json2schema

type group struct {
	schema
}

type field struct {
	schema
}

type ArrayGroupSchema interface {
	SetChildrenSchema(Schema) ArrayGroupSchema
}

type arrayGroupSchema struct {
	group
	ChildrenSchema Schema `json:"children_schema"`
}

func (arrayGroupSchema arrayGroupSchema) SetChildrenSchema(schema Schema) ArrayGroupSchema {
	arrayGroupSchema.ChildrenSchema = schema
	return arrayGroupSchema
}

type ObjectGroupSchema interface {
	AppendChild(...Schema) ObjectGroupSchema
}

type objectGroupSchema struct {
	group
	Children []Schema `json:"children"`
}

func (objectGroupSchema objectGroupSchema) AppendChild(data ...Schema) ObjectGroupSchema {
	objectGroupSchema.Children = append(objectGroupSchema.Children, data...)
	return objectGroupSchema
}

type textFieldSchema struct {
	field
}

type optionFieldSchema struct {
	field
}

type selectFieldSchema struct {
	field
	Options []optionFieldSchema `json:"options"`
}