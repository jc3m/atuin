package events

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jc3m/atuin/db/tubul"
)

type TubulEvents struct {
	Queries *tubul.Queries
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
