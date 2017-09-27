package json2schema

type TypeNeededSchema interface {
	Type(string) NameNeededSchema
}

type NameNeededSchema interface {
	Name(string) SchemaFactory
}

func Kind(kind string) TypeNeededSchema {
	var schemaFactory = new(schemaFactory)
	schemaFactory.Kind = kind
	return schemaFactory
}

func (schemaFactory *schemaFactory) Type(t string) NameNeededSchema {
	schemaFactory.schema.Type = t
	return schemaFactory
}

func (schemaFactory *schemaFactory) Name(name string) SchemaFactory {
	schemaFactory.schema.Name = name
	return schemaFactory
}