package json2schema

import (
	"testing"
	"encoding/json"
	"fmt"
	"hami/ums/base/log"
)

func TestSchema(t *testing.T) {
	//var ag = ArrayGroup(VarRoot)
	//var ag = Field("asd",FieldTypeText)
	/*var ag = ObjectGroup(VarRoot)
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))*/


	var ag = Kind(KindField).Type(FieldTypeSelect).Name(VarRoot).Make()
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))

	log.Logf(log.WithoutColor,"string(js): type: '%T' value: %v",ag,string(js))
}