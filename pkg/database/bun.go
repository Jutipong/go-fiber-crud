package database

// import (
// 	"database/sql"
// 	"time"

// 	_ "github.com/denisenkom/go-mssqldb"
// 	"github.com/uptrace/bun"
// 	"github.com/uptrace/bun/dialect/mssqldialect"
// )

// type Customer2 struct {
// 	bun.BaseModel `bun:"table:CustomerX,alias:u"`

// 	Id          string     `bun:"Id,pk"`
// 	Name        string     `bun:"Name"`
// 	Email       string     `bun:"Email"`
// 	Age         int        `bun:"Age"`
// 	CreatedDate *time.Time `bun:"CreatedDate"`
// 	IsActive    bool       `bun:"IsActive"`
// }

// func InitDbBun() *bun.DB {
// 	configDb := getConfigDb()
// 	// sqldb, err := sql.Open(sql., configDb+"&charset=utf8mb4&parseTime=True&loc=Local")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	sqldb, err := sql.Open("sqlserver", configDb)
// 	if err != nil {
// 		panic(err)
// 	}

// 	db := bun.NewDB(sqldb, mssqldialect.New())

// 	return db
// }
