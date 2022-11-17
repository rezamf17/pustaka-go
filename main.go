package main

import (
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
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	// bookRequest := book.BookRequest{
	// 	Title: "Gundam",
	// 	Price: "90000",
	// }

	// bookService.Create(bookRequest)
	// dbRepository := book.NewRepository(db)

	// book, err := dbRepository.FindByID(2)

	// fmt.Println("Title :", book.Title)
	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// bookRepository := book.NewRepository(db)
	// book := book.Book{
	// 	Title:       "Memburu Ikan Paus",
	// 	Description: "All about Fish",
	// 	Price:       15000,
	// 	Rating:      3,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)

	router := gin.Default()
	v1 := router.Group("/v1")
	router.GET("/books", bookHandler.GetAllBooks)
	router.GET("/hello", bookHandler.HelloWorld)
	router.GET("/books/:id", bookHandler.GetBookId)
	router.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run(":8888")
}
