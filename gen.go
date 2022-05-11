package main

import (
	"bach-reconcile-non-swift/pkg/config"
	"bach-reconcile-non-swift/pkg/database"

	"gorm.io/gen"
)

func init() {
	config.Init()
	database.Init()
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/ef",
		Mode:         gen.WithoutContext,
		// generate model global configuration
		FieldNullable:     true, // generate pointer when field is nullable
		FieldCoverable:    true, // generate pointer when field has default value
		FieldWithIndexTag: true, // generate with gorm index tag
		FieldWithTypeTag:  true, // generate with gorm column type tag
	})

	g.UseDB(database.DBConn)
	g.ApplyBasic(g.GenerateModel("MS_MailConfig"))
	g.Execute()
}
