package main

import (
	"fiber-crud/pkg/config"

	"gorm.io/gen"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./dal/query",
		ModelPkgPath:      "./dal/model",
		Mode:              gen.WithoutContext,
		FieldNullable:     true, // generate pointer when field is nullable
		FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})
	// db, _ := gorm.Open(sqlserver.Open("sqlserver://sa:p@ssw0rd@Localhost?database=GolangDemo&charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(config.Db())
	// card := g.GenerateModel("credit_cards")
	// customer := g.GenerateModel("customers", gen.FieldRelate(field.HasMany, "CreditCards", b,
	// 	&field.RelateConfig{
	// 		// RelateSlice: true,
	// 		GORMTag: "foreignKey:CustomerRefer",
	// 	}),
	// )

	g.WithDataTypeMap(dataMap)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

var dataMap = map[string]func(detailType string) (dataType string){
	"float64": func(detailType string) (dataType string) { return "float64" },
	// "json":  func(string) string { return "json.RawMessage" },
}

type RelateConfig struct {
	// specify field's type
	RelatePointer      bool // ex: CreditCard  *CreditCard
	RelateSlice        bool // ex: CreditCards []CreditCard
	RelateSlicePointer bool // ex: CreditCards []*CreditCard

	JSONTag      string // related field's JSON tag
	GORMTag      string // related field's GORM tag
	NewTag       string // related field's new tag
	OverwriteTag string // related field's tag
}
