package json2schema

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestSchema(t *testing.T) {
	//var ag = ArrayGroup(VarRoot)
	//var ag = Field("asd",FieldTypeText)
	/*var ag = ObjectGroup(VarRoot)
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))*/


	var ag = Kind(KindField).Type(FieldTypeSelect).SetName(VarRoot).Out()
	var js, _ = json.Marshal(ag)
	fmt.Println(string(js))
}