package handler

import (
	"fmt"
	"net/http"
	"pustaka-golang/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Introduce(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Reza",
		"bio":  "Front End Web Developer at PT. Sarana Pactindo",
	})
}

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Reza",
		"bio":  "Hello World",
	})
}

func PostBooksHandler(ctx *gin.Context) {
	var postBooks book.BookRequest

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

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}
