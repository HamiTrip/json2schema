package json2schema

import (
	"encoding/json"
	"hami/ums/base/log"
)

type Schema interface {
	Out() map[string]interface{}
}
type nameNeededSchema interface {
	SetName(string) Schema
}

type inFactorySchema interface {
	SetKind(string)
	SetType(string)
}

type schema struct {
	Attributes basicAttributes `json:"attributes"`
	Validations `json:"validations"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type basicAttributes struct {
	Title       struct {
			    language
		    } `json:"title"`
	Placeholder struct {
			    language
		    } `json:"placeholder"`
}

type language struct {
	En string `json:"en"`
}

type Validations []string

type group struct {
	schema
}

type field struct {
	schema
}

func (schema *schema) SetType(typeName string) {
	schema.Type = typeName
}

func (schema *schema) SetKind(kind string) {
	schema.Kind = kind
}

func (schema *schema) SetName(name string) Schema {
	schema.Name = name
	return schema
}

func (schema *schema) Out() map[string]interface{} {
	var js, _ = json.Marshal(schema)
	var res interface{}
	json.Unmarshal(js, res)
	log.Warning("schema:",schema)
	log.Warning("js:",string(js))
	log.Warning("res:",res)
	//return res.(map[string]interface{})
	return map[string]interface{}{}
}
