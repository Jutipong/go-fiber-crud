package main

import (
	"context"
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var dataLength = 1000

func main() {
	config.InitialConfig()
	database.InitialDB()
	//
	BunTest()
	GormTest()
}

func GormTest() {
	data := []database.Customer{}
	for i := 0; i < dataLength; i++ {
		data = append(data, database.Customer{
			Id:          uuid.NewString(),
			Name:        fmt.Sprintf("Name: %v", i),
			Email:       fmt.Sprintf("Email: %v", i),
			Age:         i,
			CreatedDate: nil,
			IsActive:    true,
		})
	}

	db := database.Db()

	start := time.Now()

	if err := db.CreateInBatches(&data, 300).Error; err != nil {
		panic(err)
	}

	end := time.Since(start)
	fmt.Println(fmt.Sprintf("Gorm => insert data: %v record time: %.2f:%.2f:%v", len(data), end.Hours(), end.Minutes(), end.Milliseconds()))
}

func BunTest() {
	data := []database.CustomerBun{}
	for i := 0; i < dataLength; i++ {
		data = append(data, database.CustomerBun{
			Id:          uuid.NewString(),
			Name:        fmt.Sprintf("Name: %v", i),
			Email:       fmt.Sprintf("Email: %v", i),
			Age:         i,
			CreatedDate: nil,
			IsActive:    true,
		})
	}

	db := database.InitDbBun()
	len := len(data)

	start := time.Now()

	_, err := db.NewInsert().Model(&data).Exec(context.Background())
	if err != nil {
		panic(err)
	}

	end := time.Since(start)
	fmt.Println(fmt.Sprintf("Bun => insert data: %v record time: %.2f:%.2f:%v", len, end.Hours(), end.Minutes(), end.Milliseconds()))
}
