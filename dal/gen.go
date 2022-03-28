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
		OutPath: "./dal/query",
		Mode:    gen.WithoutContext,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})
	// db, _ := gorm.Open(sqlserver.Open("sqlserver://sa:p@ssw0rd@Localhost?database=GolangDemo&charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(config.Db())
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
