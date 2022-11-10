package main

import (
	"fmt"
	"log"
	"pustaka-golang/book"
	"pustaka-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=reza-golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error")
	}
	db.AutoMigrate(&book.Book{})

	// book := book.Book{}
	// book.Title = "Ikan Manis"
	// book.Price = 90000
	// book.Discount = 40
	// book.Description = "Hidup sudah sulit"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Print("Error Insert")
	// }

	// var books []book.Book

	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	fmt.Println("Data not found")
	// }

	// for _, item := range books {
	// 	fmt.Println("Title : ", item.Title)
	// 	fmt.Println("Book Object : %v", item)
	// }
	var book book.Book
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("Data not found")
	}
	book.Title = "Tiger Wong Son of Wong"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("Data change failed")
	}

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("Data delete failed")
	}

	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/", handler.Introduce)
	router.GET("/hello", handler.HelloWorld)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
