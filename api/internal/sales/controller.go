package sales

import (
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
    "strconv"
)

// Products
func AllProducts(c *gin.Context) {
    records := []Product{}
	if err := util.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func GetProduct(c *gin.Context) {
    record := Product{}
	if err := util.DB.First(&record, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func AddProduct(c *gin.Context) {
    record := Product{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

func UpdateProduct(c *gin.Context) {
    record := Product{}
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

func DeleteProduct(c *gin.Context) {
    record := Product{}
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
    record.ID = uint(id)
	util.DB.Delete(&record)
	c.JSON(http.StatusOK, record)
}

// Sales
func AllSales(c *gin.Context) {
    records := []Sale{}
	if err := util.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func GetSale(c *gin.Context) {
    record := Sale{}
	if err := util.DB.First(&record, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func AddSale(c *gin.Context) {
    record := Sale{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	util.DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

func UpdateSale(c *gin.Context) {
    record := Sale{}
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

func DeleteSale(c *gin.Context) {
    record := Sale{}
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
    record.ID = uint(id)
	util.DB.Delete(&record)
	c.JSON(http.StatusOK, record)
}
