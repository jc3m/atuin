package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jc3m/atuin/db/tubul"
	"github.com/jc3m/atuin/tubul/events"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=atuin password=password dbname=atuin host=localhost port=5200 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	queries := tubul.New(db)

	router := gin.Default()

	te := events.TubulEvents{Queries: queries}
	events := router.Group("/events")
	{
		events.GET("/", te.Get)
		events.POST("/", te.Post)
	}

	router.Run("localhost:8080")
}
