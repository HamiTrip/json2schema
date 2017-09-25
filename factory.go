package json2schema

import (
	"reflect"
)

type SchemaFactory interface {
	Make() Schema
}

type schemaFactory struct {
	schema
}

type typeName string

type schemaType reflect.Type

type schemaKinds map[typeName]schemaType

type schemas map[string]schemaKinds

var allSchemas = schemas{
	KindGroup: schemaKinds{
		TypeObject:reflect.TypeOf(objectGroupSchema{}),
		TypeArray:reflect.TypeOf(arrayGroupSchema{}),
	},
	KindField: schemaKinds{
		FieldTypeText:reflect.TypeOf(textFieldSchema{}),
		FieldTypeSelect:reflect.TypeOf(selectFieldSchema{}),
		FieldTypeOption:reflect.TypeOf(optionFieldSchema{}),
	},
}

func (schemaFactory *schemaFactory) Make() Schema {
	if sType, ok := allSchemas[schemaFactory.Kind][typeName(schemaFactory.schema.Type)]; !ok {
		panic("invalid type or kind!")
	} else {
		var rSchema = reflect.Indirect(reflect.New(sType))

		rSchema.FieldByName("Kind").Set(reflect.ValueOf(schemaFactory.Kind))
		rSchema.FieldByName("Type").Set(reflect.ValueOf(schemaFactory.schema.Type))
		rSchema.FieldByName("Name").Set(reflect.ValueOf(schemaFactory.schema.Name))
		rSchema.FieldByName("Validations").Set(reflect.ValueOf([]string{"required"}))



		return (rSchema).Interface().(Schema)
	}
}