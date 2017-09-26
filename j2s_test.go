package json2schema

import (
	"testing"
	"encoding/json"
	"fmt"
	"hami/ums/base/log"
	"io/ioutil"
)

func TestMakeSchema(t *testing.T) {
	//var ag = ArrayGroup(VarRoot)
	//var ag = Field("asd",FieldTypeText)
	/*var ag = ObjectGroup(VarRoot)
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))*/


	//var ag = Kind(KindField).Type(FieldTypeSelect).Name(VarRoot).Make()

	data, e := ioutil.ReadFile("/home/sabloger/code/go/sandbox/src/github.com/hamitrip/json2schema/json_sample/b2x_final.json")
	if e != nil {
		t.Errorf("File error: %v\n", e)
	}
	/*schema, e := ioutil.ReadFile("/home/sabloger/code/go/sandbox/src/hami/ums/test/test_recources/schema-2.json")
	if e != nil {
		t.Errorf("File error: %v\n", e)
	}*/
	//var schemaMap map[string]interface{}
	var dataMap interface{}

	json.Unmarshal(data, &dataMap)
	//json.Unmarshal(schema, &schemaMap)

	var ag = MakeSchema(dataMap)
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))

	log.Logf(log.WithoutColor,"string(js): type: '%T' value: %v",ag,string(js))
}