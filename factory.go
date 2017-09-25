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
		TypeObject:reflect.TypeOf(objectGroup{}),
		TypeArray:reflect.TypeOf(arrayGroup{}),
	},
	KindField: schemaKinds{
		FieldTypeText:reflect.TypeOf(textField{}),
		FieldTypeSelect:reflect.TypeOf(selectField{}),
		FieldTypeOption:reflect.TypeOf(optionField{}),
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

		return (rSchema).Interface().(Schema)
	}
}