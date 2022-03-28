package main

import (
	"encoding/json"
	"fiber-crud/dal/model"
	"fiber-crud/dal/query"
	"fiber-crud/middleware"
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fiber-crud/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shopspring/decimal"
	"gorm.io/gen"
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

	// Middleware
	app.Use(middleware.Logger)
	// app.Use(middleware.Authorization())
	// genSQL()

	q := query.Use(config.Db())
	address := q.Address
	user := q.User
	// u, _ := user.Preload(field.Associations).Find()
	// fmt.Println(u)

	// a, _ := q.Address.Preload(field.Associations).Find()
	// fmt.Println(a)
	var result model.User
	user.Select(user.ALL).
		LeftJoin(address, address.AddressID.EqCol(user.AddressID)).Scan(&result)
	str, _ := json.Marshal(&result)
	fmt.Println(string(str))
	// lat, err := decimal.NewFromString(string(u.Lat))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// long, err := decimal.NewFromString(string(u.Long))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(lat.Add(long))
	// var tt []uint8

	// tt = decimal.NewFromString("")

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

func genSQL() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
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
