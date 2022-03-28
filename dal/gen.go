package main

import (
	"fiber-crud/pkg/config"
	"strings"

	"gorm.io/gen"
)

func init() {

}

// dataMap mapping relationship
var dataMap = map[string]func(detailType string) (dataType string){
	// int mapping
	"int":   func(detailType string) (dataType string) { return "int32" },
	"int32": func(detailType string) (dataType string) { return "int32" },
	"int64": func(detailType string) (dataType string) { return "int32" },

	// bool mapping
	"tinyint": func(detailType string) (dataType string) {
		if strings.HasPrefix(detailType, "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/model",
		Mode:         gen.WithoutContext,

		// generate model global configuration
		FieldNullable:     true, // generate pointer when field is nullable
		FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})

	g.UseDB(config.Db())

	// specify diy mapping relationship
	g.WithDataTypeMap(dataMap)

	// generate all field with json tag end with "_example"
	// g.WithJSONTagNameStrategy(func(c string) string { return c + "_example" })

	mytable := g.GenerateModel("Address")
	g.ApplyBasic(mytable)
	// g.ApplyBasic(g.GenerateAllTable()...) // generate all table in db server

	g.Execute()
}

// var dataMap = map[string]func(detailType string) (dataType string){
// 	"uint8 ": func(detailType string) (dataType string) { return "float64" },
// }
