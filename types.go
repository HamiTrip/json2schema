package json2schema

type arrayGroup struct {
	group
	ChildrenSchema Schema `json:"children_schema"`
}

type objectGroup struct {
	group
	Children []Schema `json:"children"`
}

type textField struct {
	field
}

type optionField struct {
	field
}

type selectField struct {
	field
	Options []optionField `json:"options"`
}

func (selectField selectField) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"salam":"khubi?",
	}
}
