package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WeighIn struct {

	ProduceID		int		'json:"produce_id"'
	GrossWeight		float64		'json:"gross_weight"'
	ToMarketWeight		float64		'json:"to_market_weight"'
	SoldToStudents		float64		'json:"sold_to_students"'
	ValueAdded		float64		'json:"value_added"'
	Cull		float64		'json:"cull"'
	Timestamp		time.Time		'json:"timestamp"'
}

func main() {
	r := gin.Default()

	r.POST("/weigh-in", func(c *gin.Context) {
		var w WeighIn
		if err := c.ShouldBindJSON(&w); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// insert into DB here

		c.JSON(http.StatusOK, gin.H{"status": "data received", "data": w})
	})
	r.Run(":8080") //listen and serve on 0.0.0.0:8080
}

var storedWeighIns []WeighIn // testing

r.GET("/weigh-ins", func(c *gin.Context) {
    c.JSON(http.StatusOK, storedWeighIns)
})
