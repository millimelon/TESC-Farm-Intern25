package labor

import (
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
  "golang.org/x/crypto/bcrypt"
)

func hashANum(anum string) (string, error) {
  bytenum := []byte(anum)
  hash, err := bcrypt.GenerateFromPassword(bytepass, bcrypt.MinCost)
  if err != nil {
    return "", err
  }
  return string(hash), err
}

// Hours
func AllHours(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Preload("Worker").Preload("Harvest").Preload("Harvest.Crop").Preload("Process").Preload("Process.Harvest").Preload("Process.Harvest.Crop").Preload("Process.Product").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}
func HarvestingHours(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Preload("Worker").Preload("Harvest").Preload("Harvest.Crop").Where("harvest_id NOT NULL").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func ProcessingHours(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Preload("Worker").Preload("Process").Where("process_id NOT NULL").Find(&records).Error; err != nil {
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

func GetWorkerHours(c *gin.Context) {
	w := Worker{}
	if err := util.DB.First(&w, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	records := []Hours{}
	if err := util.DB.Find(&records, Hours{WorkerID: w.ID}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, h := range records {
		h.Worker = &w
	}
	c.JSON(http.StatusOK, records)
}

func AddWorker(c *gin.Context) {
	record := Worker{}
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  //Hashes A number
  hashed_a_num, err := hashANum(record.Barcode)
  //checks if A num is hashed
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
    return
  }
  record.Barcode = hashed_a_num
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

func AddTask(c *gin.Context) {
  record := Task{}
  if err := c.ShouldBindJSON(&record); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
    return
  }
  util.DB.Create(&record)
  c.JSON(http.StatusOK, record)
}

func GetTask(c *gin.Context) {
  record := Task{}
	if err := util.DB.First(&w, c.Param("id")).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, record)
}

func GetAllTasks(c *gin.Context) {
  record := []Tasks{}
	if err := util.DB.Find(&records).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, records)
}

func UpdateTask(c *gin.Context) {
  record := Task{}
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
