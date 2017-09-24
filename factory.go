package json2schema

import (
	"reflect"
)

type SchemaFactory interface {
	Type(string) nameNeededSchema
}

type schemaFactory struct {
	schemaKinds
	kind string
}

func (schemaFactory *schemaFactory)Type(tName string) nameNeededSchema {
	if s, ok := schemaFactory.schemaKinds[typeName(tName)]; ok {
		var res = reflect.New(s).Interface().(inFactorySchema)
		res.SetKind(schemaFactory.kind)
		res.SetType(tName)
		return res.(nameNeededSchema)
	} else {
		panic("invalid schema")
	}
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

func Kind(kind string) SchemaFactory {
	/*var schemaFactory = new(schemaFactory)
	schemaFactory.schemaKind = kind*/
	return &schemaFactory{
		schemaKinds: allSchemas[kind],
		kind:kind,
	}
}

/*
func ArrayGroup(name string) *arrayGroup {
	var schema = new(arrayGroup)
	schema.Name = name
	schema.Type = TypeArray
	schema.Kind = KindGroup

	return schema
}

func ObjectGroup(name string) *arrayGroup {
	var schema = new(arrayGroup)
	schema.Name = name
	schema.Type = TypeObject
	schema.Kind = KindGroup
	schema.Validations = []string{}
	return schema
}

func Field(name, t string) *arrayGroup {
	var schema = new(arrayGroup)
	schema.Name = name
	schema.Type = t
	schema.Kind = KindField
	schema.Validations = []string{"required"}
	return schema
}
*/

