package json2schema

//TODO:: Make it interactive!! Ask Titles and validations per each field!!!  Of course it only can used in cmd mode in this case!!

type SchemaGenerator interface {
	JSON(interface{}) SchemaGenerator
	Schema()
}

type schemaGenerator struct {

}