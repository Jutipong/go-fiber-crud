package main

import (
	"encoding/json"
	"fiber-crud/pkg/config"

	"gorm.io/gen"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/schema",
		Mode:         gen.WithoutContext,
		// generate model global configuration
		FieldNullable:     true, // generate pointer when field is nullable
		FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})

	g.UseDB(config.Db())

	byt, err := json.Marshal(g.GenerateAllTable())
	if err != nil {
		panic("All Talble json Marshal fail.")
	}

	var tableNames tableNameAll
	err = json.Unmarshal(byt, &tableNames)
	if err != nil {
		panic("All Talble json Unmarshal fail.")
	}

	//change type (*[]uint8 and other) to decimal
	// for _, table := range tableNames {
	// 	currentTable := g.GenerateModel(table.TableName)
	// 	for _, i := range currentTable.Fields {
	// 		if i.Type == "*[]uint8" {
	// 			i.Type = "decimal.Decimal"
	// 		}
	// 	}
	// 	g.ApplyBasic(currentTable)
	// }

	// g.ApplyBasic(mytable)
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}

type tableNameAll []struct {
	// GenBaseStruct bool   `json:"GenBaseStruct"`
	// FileName      string `json:"FileName"`
	// S             string `json:"S"`
	// NewStructName string `json:"NewStructName"`
	// StructName    string `json:"StructName"`
	TableName string `json:"TableName"`
	// StructInfo    struct {
	// 	PkgPath   string `json:"PkgPath"`
	// 	Package   string `json:"Package"`
	// 	Name      string `json:"Name"`
	// 	Type      string `json:"Type"`
	// 	IsArray   bool   `json:"IsArray"`
	// 	IsPointer bool   `json:"IsPointer"`
	// } `json:"StructInfo"`
	// Fields []struct {
	// 	Name             string      `json:"Name"`
	// 	Type             string      `json:"Type"`
	// 	ColumnName       string      `json:"ColumnName"`
	// 	ColumnComment    string      `json:"ColumnComment"`
	// 	MultilineComment bool        `json:"MultilineComment"`
	// 	JSONTag          string      `json:"JSONTag"`
	// 	GORMTag          string      `json:"GORMTag"`
	// 	NewTag           string      `json:"NewTag"`
	// 	OverwriteTag     string      `json:"OverwriteTag"`
	// 	Relation         interface{} `json:"Relation"`
	// } `json:"Fields"`
	// Source         int         `json:"Source"`
	// ImportPkgPaths interface{} `json:"ImportPkgPaths"`
}
