package main

import (
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

func main() {
	//step up router for handling endpoints
	//create router to handle different routes
	router := gin.Default()
	router.GET("/books", getBooks)
	//run webserver
	router.Run("localhost:8080")
}
