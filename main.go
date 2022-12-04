package main

import (
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	config.InitialConfig()
	database.InitialDB()
	//
	GormTest()
}

func GormTest() {
	data := MockData()
	db := database.Db()

	start := time.Now()

	if err := db.CreateInBatches(&data, 300).Error; err != nil {
		panic(err)
	}

	end := time.Since(start)
	fmt.Println(fmt.Sprintf("insert data: %v record time: %.2f:%.2f:%v", len(data), end.Hours(), end.Minutes(), end.Milliseconds()))
}

// func BunTest() {
// 	data := MockData()
// 	db := database.InitDbBun()

// 	res, err := db.NewInsert().Model(&data).Exec(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res)

// 	start := time.Now()

// 	end := time.Since(start)
// 	fmt.Println(fmt.Sprintf("insert data: %v record time: %.2f:%.2f:%v", len(data), end.Hours(), end.Minutes(), end.Milliseconds()))
// }

func MockData() []database.Customer {
	result := []database.Customer{}
	for i := 0; i < 100; i++ {
		result = append(result, database.Customer{
			Id:          uuid.NewString(),
			Name:        fmt.Sprintf("Name: %v", i),
			Email:       fmt.Sprintf("Email: %v", i),
			Age:         i,
			CreatedDate: nil,
			IsActive:    true,
		})
	}
	return result
}
