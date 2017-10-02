package json2schema

import (
	"encoding/json"
)

/*
Schema interface
 */
type Schema interface {
	ToMap() map[string]interface{}
}

type schema struct {
	Attributes  basicAttributes `json:"attributes"`
	Validations `json:"validations"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type basicAttributes struct {
	Title struct {
		language
	} `json:"title"`
	Placeholder struct {
		language
	} `json:"placeholder"`
}

type language struct {
	En string `json:"en"`
}

/*
Validations is list of validator conditions
 */
type Validations []string

func (schema schema) ToMap() map[string]interface{} {
	var js, err = json.Marshal(schema)
	if err != nil {
		panic(err)
	}
	var res interface{}
	err = json.Unmarshal(js, &res)
	if err != nil {
		panic(err)
	}
	return res.(map[string]interface{})
	// shayad beshe akhare kar ba Cast ya Convert be type asly dataye kamelo bekeshim birun!!
}
