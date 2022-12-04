package database

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

type CustomerBun struct {
	bun.BaseModel `bun:"table:Customer_Bun,alias:u"`

	Id          string     `bun:"Id,pk"`
	Name        string     `bun:"Name"`
	Email       string     `bun:"Email"`
	Age         int        `bun:"Age"`
	CreatedDate *time.Time `bun:"CreatedDate"`
	IsActive    bool       `bun:"IsActive"`
}

func InitDbBun() *bun.DB {
	configDb := getConfigDb()
	sqldb, err := sql.Open("sqlserver", configDb)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mssqldialect.New())

	return db
}
