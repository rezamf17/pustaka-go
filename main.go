package main

import (
	"pustaka-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/", handler.Introduce)
	router.GET("/hello", handler.HelloWorld)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
