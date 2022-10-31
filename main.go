package main

import (
	"fmt"
	"log"
	"pustaka-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=reza-golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error")
	}
	fmt.Println("connection success")

	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/", handler.Introduce)
	router.GET("/hello", handler.HelloWorld)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
