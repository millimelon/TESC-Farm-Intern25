package labor

import (
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func hashANum(anum string) (string, error) {
	bytenum := []byte(anum)
	hash, err := bcrypt.GenerateFromPassword(bytenum, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// Hours
func AllHours(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Preload("Worker").Preload("Task").Preload("Task.Area").Preload("Task.Planting").Preload("Task.Planting.Crop").Preload("Task.Planting.Bed").Preload("Task.Planting.Bed.Area").Preload("Task.Harvest").Preload("Task.Harvest.Crop").Preload("Task.Harvest.Bed").Preload("Task.Harvest.Bed.Area").Preload("Task.Process").Preload("Task.Process.Harvest").Preload("Task.Process.Harvest.Crop").Preload("Task.Process.Harvest.Bed").Preload("Task.Process.Harvest.Bed.Area").Preload("Task.Process.Product").Find(&records).Error; err != nil {
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

func GetWorking(c *gin.Context) {
	records := []Hours{}
	if err := util.DB.Preload("Worker").Preload("Task").Preload("Task.Area").Preload("Task.Planting").Preload("Task.Planting.Crop").Preload("Task.Planting.Bed").Preload("Task.Planting.Bed.Area").Preload("Task.Harvest").Preload("Task.Harvest.Crop").Preload("Task.Harvest.Bed").Preload("Task.Harvest.Bed.Area").Preload("Task.Process").Preload("Task.Process.Harvest").Preload("Task.Process.Harvest.Crop").Preload("Task.Process.Harvest.Bed").Preload("Task.Process.Harvest.Bed.Area").Preload("Task.Process.Product").Where("Duration = ?", 0).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
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

func AddPunch(c *gin.Context) {
  type ScanPunch struct {
    Barcode string `json:"barcode"`
    TaskID  int    `json:"task,omitempty"`
  }
  punch := ScanPunch{}
	if err := c.ShouldBindJSON(&punch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	anum, err := hashANum(record.Barcode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	last := Hours{}
  if err := util.DB.Preload("Worker").Where("Worker.Barcode = ?", anum).Order("Start desc").First(&last); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  if last.Duration == 0 {
    last.Duration = time.Now().Sub(last.Start)
    util.DB.Save(&last)
  }
  if taskID == 0 {
    c.JSON(http.StatusOK, last)
    return
  }
  record := Hours{}
  record.Start = time.Now()
  record.Duration = 0
  record.WorkerID = last.Worker.ID
  record.TaskID = taskID
	util.DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

{"barcode": "A# here", "task": "task id here"}

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
	if err := util.DB.Preload("Area").Preload("Planting").Preload("Planting.Crop").Preload("Planting.Bed").Preload("Planting.Bed.Area").Preload("Harvest").Preload("Harvest.Crop").Preload("Harvest.Bed").Preload("Harvest.Bed.Area").Preload("Process").Preload("Process.Harvest").Preload("Process.Harvest.Crop").Preload("Process.Harvest.Bed").Preload("Process.Harvest.Bed.Area").Preload("Process.Product").First(&record, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}

func GetAllTasks(c *gin.Context) {
	records := []Task{}
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
