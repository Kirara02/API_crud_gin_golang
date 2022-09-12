package main

import (
	"log"
	"web-api-gin-tutorial/book"
	"web-api-gin-tutorial/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "root:@tcp(127.0.0.1:3306)/web-api-gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)
	// bookRequest := book.BookRequest{
	// 	Title: "jab",
	// }
	// bookService.Create(bookRequest)

	v1 := router.Group("/v1")
	v1.POST("/books", bookHandler.CreateBooks)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
