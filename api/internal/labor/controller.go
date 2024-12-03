package labor

import (
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Hours
func AllHours(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func GetHours(c *gin.Context) {
	record := Hours{}
	if err := util.DB.First(&record, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func AddHours(c *gin.Context) {
	record := Hours{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

func UpdateHours(c *gin.Context) {
	record := Hours{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record.ID = uint(id)
	util.DB.Save(&record)
	c.JSON(http.StatusOK, record)
}

func DeleteHours(c *gin.Context) {
	record := Hours{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record.ID = uint(id)
	util.DB.Delete(&record)
	c.JSON(http.StatusOK, record)
}

// Workers
func AllWorkers(c *gin.Context) {
	records := []Worker{}
	if err := util.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func GetWorker(c *gin.Context) {
	record := Worker{}
	if err := util.DB.First(&record, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func AddWorker(c *gin.Context) {
	record := Worker{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

func UpdateWorker(c *gin.Context) {
	record := Worker{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record.ID = uint(id)
	util.DB.Save(&record)
	c.JSON(http.StatusOK, record)
}

func DeleteWorker(c *gin.Context) {
	record := Worker{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record.ID = uint(id)
	util.DB.Delete(&record)
	c.JSON(http.StatusOK, record)
}
