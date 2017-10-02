package json2schema

//TODO:: Make it interactive!! Ask Titles and validations per each field!!!
// ->Of course it only can used in cmd mode in this case!!

/*
MakeSchema is base starter method
 */
func MakeSchema(jsonData interface{}) Schema {
	//panic(fmt.Sprintf("Excepted types was 'map[string]interface{}' and '[]interface{}' got: %T", jsonData))
	var root = hub(VarRoot, jsonData)
	return root
}

func hub(name string, data interface{}) Schema {
	switch data.(type) {
	case map[string]interface{}:
		return makeObject(name, data.(map[string]interface{}))
	case []interface{}:
		return makeArray(name, data.([]interface{}))
	default:
		return makeField(name, data)
	}
}

func makeObject(name string, data map[string]interface{}) Schema {
	var schema = Kind(KindGroup).Type(TypeObject).Name(name).Make()
	for key, value := range data {
		schema = schema.(ObjectGroupSchema).AppendChild(hub(key, value)).(Schema)
	}
	return schema
}

func makeArray(name string, data []interface{}) Schema {
	var schema = Kind(KindGroup).Type(TypeArray).Name(name).Make()
	if(len(data) > 0) {
		schema = schema.(ArrayGroupSchema).SetChildrenSchema(hub(VarArrayRoot, data[0])).(Schema)
	}
	return schema
}

func makeField(name string, data interface{}) Schema {
	//TODO:: In kararo tuye khode builder kone! inja faghat callesh kone!?
	switch data.(type) {
	case int, int8, int16, int32, int64,
	uint, uint8, uint16, uint32, uint64,
	float32, float64:
		return Kind(KindField).Type(FieldTypeText).Name(name).Make()
	default:
		return Kind(KindField).Type(FieldTypeText).Name(name).Make()
	}
}