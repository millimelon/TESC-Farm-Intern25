package main

import (
	"github.com/absentbird/TESC-Farm/internal/harvest"
	"github.com/absentbird/TESC-Farm/internal/labor"
	"github.com/absentbird/TESC-Farm/internal/sales"
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Initialize the config, database, and router
	conf := util.Config
	conf.Load()
	util.ConnectDB(conf.DBConn)
	r := gin.Default()

	// Labor endpoints
	r.GET("/hours", labor.AllHours)
	r.GET("/hours/:id", labor.WorkerHours)
	r.POST("/hours/:id/update", labor.AdjustHours)
	r.POST("/hours/new", labor.AddHours)
	r.GET("/workers", labor.AllWorkers)
	r.GET("/worker/:id", labor.GetWorker)
	r.POST("/worker/:id/update", labor.UpdateWorker)
	r.POST("/worker/new", labor.AddWorker)

	// Harvest endpoints
	r.GET("/crops", harvest.AllCrops)
	r.GET("/crop/:id", harvest.GetCrop)
	r.POST("/product/:id/update", harvest.UpdateProduct)
	r.POST("/product/new", harvest.NewProduct)
	r.GET("/harvests", harvest.AllHarvests)
	r.GET("/harvest/:id", harvest.GetHarvest)
	r.POST("/harvest/:id/update", harvest.UpdateHarvest)
	r.POST("/harvest/new", harvest.AddHarvest)

	// Sales Endpoints
	r.GET("/sales", sales.AllSales)
	r.GET("/sale/:id", sales.GetSale)
	r.POST("/sale/:id/update", sales.UpdateSale)
	r.POST("/sale/new", sales.AddSale)

	err := r.Run("127.0.0.1:" + conf.Port)
	util.Check(err, "Error starting API")
}
