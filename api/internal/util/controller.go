package util

import (
	"github.com/gin-gonic/gin"
    "net/http"
)

func SusSurvey(c *gin.Context) {
    type SUS struct {
        answers []int `json:"answers"`
    }
    input := SUS{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing answers"})
        return
    }
    
}
