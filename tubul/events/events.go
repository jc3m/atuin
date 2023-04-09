package events

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jc3m/atuin/db/tubul"
)

type TubulEvents struct {
	Queries *tubul.Queries
}

type errorMessageStruct struct {
	Message string `json:"message"`
}

type createEventParams struct {
	Title     string `json:"title"`
	EventDate string `json:"event_date"`
}

func (t *TubulEvents) Get(c *gin.Context) {
	events, err := t.Queries.ListEvents(c)

	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func (t *TubulEvents) Post(c *gin.Context) {
	var params createEventParams

	if err := c.BindJSON(&params); err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	eventDate, err := time.Parse("2006-01-02", params.EventDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorMessageStruct{"Event date must be in format YYYY-MM-DD"})
	}
	dbParams := tubul.CreateEventParams{
		Title:     params.Title,
		EventDate: eventDate,
	}
	event, err := t.Queries.CreateEvent(c, dbParams)

	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, event)
}
