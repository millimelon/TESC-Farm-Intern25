package labor

import (
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
    "strconv"
)

func AllHours(c *gin.Context) {
    hours := []Hours{}
	if err := util.DB.Find(&hours).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hours)
}

func GetHours(c *gin.Context) {
    hours := Hours{}
	if err := util.DB.First(&hours, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hours)
}

func UpdateHours(c *gin.Context) {
    hours := Hours{}
	if err := c.ShouldBindJSON(&hours); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
    hours.ID = uint(id)
	util.DB.Save(&hours)
	c.JSON(http.StatusOK, gin.H{"data": hours})
}

func AddHours(c *gin.Context) {
    hours := Hours{}
	if err := c.ShouldBindJSON(&hours); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&hours)
	c.JSON(http.StatusOK, gin.H{"data": hours})
}

func AllWorkers(c *gin.Context) {
    workers := []Worker{}
	if err := util.DB.Find(&workers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workers)
}

func GetWorker(c *gin.Context) {
    worker := Worker{}
	if err := util.DB.First(&worker, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, worker)
}

func UpdateWorker(c *gin.Context) {
    worker := Worker{}
	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
    worker.ID = uint(id)
	util.DB.Save(&worker)
	c.JSON(http.StatusOK, gin.H{"data": worker})
}

func AddWorker(c *gin.Context) {
    worker := Worker{}
	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&worker)
	c.JSON(http.StatusOK, gin.H{"data": worker})
}
