package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	//serialize go struct to json format
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "John Doe", Quantity: 2},
	{ID: "2", Title: "The Times have Changed", Author: "Liam Jack", Quantity: 9},
	{ID: "3", Title: "The Voyager", Author: "Martin Odegwu", Quantity: 6},
}

// get request
// *gin.Context allows user to handle the request and return a response
func getBooks(c *gin.Context) {
	//give nice indented JSON format
	//books is serialized to JSON
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		//return error
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// get book by id
func getBookById(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			//if found return pointer in books struct
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

// add book
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		//pass error response message
		return
	}

	//if successfull add the newbook to books
	books = append(books, newBook)

	//return the newbook
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	//step up router for handling endpoints
	//create router to handle different routes
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	//run webserver
	router.Run("localhost:8080")
}
