package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TubulEvent struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"Date"`
}

var events = []TubulEvent{
	{ID: "1", Title: "Brandon Johnson elected Mayor of Chicago", Date: "2023-04-04"},
	{ID: "2", Title: "Silicon Valley Bank collapses", Date: "2023-03-10"},
	{ID: "3", Title: "Google lays off 12,000 employees", Date: "2023-01-20"},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func main() {
	router := gin.Default()
	router.GET("/events", getEvents)

	router.Run("localhost:8080")
}
