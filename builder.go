package json2schema

/*
TypeNeededSchema is one of builder methods
 */
type TypeNeededSchema interface {
	Type(string) nameNeededSchema
}

type nameNeededSchema interface {
	Name(string) SchemaFactory
}

/*
Kind will run builder and set kind
 */
func Kind(kind string) TypeNeededSchema {
	var schemaFactory = new(schemaFactory)
	schemaFactory.Kind = kind
	return schemaFactory
}

func (schemaFactory *schemaFactory) Type(t string) nameNeededSchema {
	schemaFactory.schema.Type = t
	return schemaFactory
}

func (schemaFactory *schemaFactory) Name(name string) SchemaFactory {
	schemaFactory.schema.Name = name
	return schemaFactory
}