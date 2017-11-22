package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	var err error
	// db, err = gorm.Open("sqlite3", "test.db")
	// db, err := gorm.Open("postgres", "user=username dbname=password sslmode=disable")

	db, err := gorm.Open("mysql", "root:root@/stars?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	fmt.Printf("test using interface\n")

	//	defer db.Close()
}
