package main

import (
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func init() {
	config.InitialConfig()
	database.InitialDB()
}

func main() {

	data := MockData()
	db := database.Db()

	start := time.Now()
	err := db.CreateInBatches(&data, 300).Error
	if err != nil {
		fmt.Println(err)
	}

	end := time.Since(start)

	// var elapsedTime = String.Format("insert data: {0} record time: {1:00}:{2:00}:{3:00}.{4:00}", customers.Count(),
	// ts.Hours, ts.Minutes, ts.Seconds,
	// ts.Milliseconds / 10);

	fmt.Println(fmt.Sprintf("insert data: %v record time: %.2f:%.2f:%v", len(data), end.Hours(), end.Minutes(), end.Milliseconds()))

}

func MockData() []database.Customer {
	// start := time.Now()
	result := []database.Customer{}
	for i := 0; i < 500000; i++ {
		result = append(result, database.Customer{
			Id:          uuid.NewString(),
			Name:        fmt.Sprintf("Name: %v", i),
			Email:       fmt.Sprintf("Email: %v", i),
			Age:         i,
			CreatedDate: nil,
			IsActive:    true,
		})
	}

	// end := time.Since(start)
	// fmt.Println(fmt.Sprintf("mock data: %v", end))

	return result
}
