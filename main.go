package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", introduce)

	router.GET("/hello", helloWorld)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8888")
}

func introduce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Reza",
		"bio":  "Front End Web Developer at PT. Sarana Pactindo",
	})
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Reza",
		"bio":  "Hello World",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type PostBooks struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,numeric"`
}

func postBooksHandler(ctx *gin.Context) {
	var postBooks PostBooks

	err := ctx.ShouldBindJSON(&postBooks)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, Condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			// ctx.JSON(http.StatusBadRequest, errorMessage)
			// return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": postBooks.Title,
		"price": postBooks.Price,
	})
}
