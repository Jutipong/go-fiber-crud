package main

import (
	"fiber-crud/middleware"
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fiber-crud/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shopspring/decimal"
)

func init() {
	config.InitialConfig()
	config.InitialDB()
}

type Result struct {
	Name    string
	Address string
}

func main() {
	decimal.MarshalJSONWithoutQuotes = true
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*"}))
	//test 1
	//test 2
	// Middleware
	app.Use(middleware.Logger)
	// app.Use(middleware.Authorization())
	// genSQL()

	// q := query.Use(config.Db())
	// // address := q.Address
	// // user := q.User
	// // // u, _ := user.Preload(field.Associations).Find()
	// // // fmt.Println(u)

	// a, _ := q.Address.Find()
	// for _, item := range a {
	// 	total := item.Lat.Add(item.Long).Add(item.Price)
	// 	fmt.Println(total)
	// }

	// Routes
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	_configEnv := config.Server()
	if _configEnv.Env_Mode == enum.ModeDebug {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

// dataMap mapping relationship
var dataMap = map[string]func(detailType string) (dataType string){
	// int mapping
	"int": func(detailType string) (dataType string) { return "int32" },

	// bool mapping
	"tinyint": func(detailType string) (dataType string) {
		if strings.HasPrefix(detailType, "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},
}

// func genSQL() {
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath:      "./dal/query",
// 		ModelPkgPath: "./dal/model",
// 		Mode:         gen.WithoutContext,
// 		// generate model global configuration
// 		FieldNullable:     true, // generate pointer when field is nullable
// 		FieldCoverable:    true, // generate pointer when field has default value
// 		FieldWithIndexTag: true, // generate with gorm index tag
// 		FieldWithTypeTag:  true, // generate with gorm column type tag
// 	})

// 	g.UseDB(config.Db())

// 	//change type (*[]uint8 and other) to decimal
// 	byt, err := json.Marshal(g.GenerateAllTable())
// 	if err != nil {
// 		panic("All Talble json Marshal fail.")
// 	}

// 	var tableNames tableNameAll
// 	err = json.Unmarshal(byt, &tableNames)
// 	if err != nil {
// 		panic("All Talble json Unmarshal fail.")
// 	}

// 	for _, table := range tableNames {
// 		currentTable := g.GenerateModel(table.TableName)
// 		for _, i := range currentTable.Fields {
// 			if i.Type == "*[]uint8" {
// 				i.Type = "decimal.Decimal"
// 			}
// 		}
// 		g.ApplyBasic(currentTable)
// 	}

// 	// g.ApplyBasic(mytable)
// 	// g.ApplyBasic(g.GenerateAllTable()...) // generate all table in db server

// 	g.Execute()
// }
