package main

import (
	"github.com/absentbird/TESC-Farm/internal/harvest"
	"github.com/absentbird/TESC-Farm/internal/labor"
	"github.com/absentbird/TESC-Farm/internal/sales"
	"github.com/absentbird/TESC-Farm/internal/util"
	"github.com/absentbird/TESC-Farm/internal/util/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	PRODCONFIG = "configs/config.yaml"
	DEVCONFIG  = "configs/config-dev.yaml"
)

func main() {
	// Initialize the config, database, and router
	conf := util.Config{}
	conf.Load(PRODCONFIG, DEVCONFIG)
	util.DB = db.ConnectDB(conf.DBConn)
	util.WorkerToken = conf.WToken
	if conf.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(cors.Default())

	// Session endpoints
	r.POST("/login", util.Login)
	r.GET("/logout", util.Logout)

	// Group for authorized endpoints
	api := r.Group("/")
	//api.Use(util.AuthMiddleware())

	// Labor endpoints
	api.GET("/hours", labor.AllHours)
	api.GET("/hours/:id", labor.GetHours)
	api.GET("/hours/working", labor.GetWorking)
	api.POST("/hours/:id/update", labor.UpdateHours)
	api.POST("/hours/:id/delete", labor.DeleteHours)
	api.POST("/hours/new", labor.AddHours)
	api.POST("/hours/punch", labor.AddPunch)
	api.GET("/tasks", labor.AllTasks)
	api.GET("/task/:id", labor.GetTask)
	api.POST("/task/:id/update", labor.UpdateTask)
	api.POST("/task/:id/delete", labor.DeleteTask)
	api.POST("/task/new", labor.AddTask)
	api.GET("/workers", labor.AllWorkers)
	api.GET("/worker/:id", labor.GetWorker)
	api.GET("/worker/:id/hours", labor.GetWorkerHours)
	api.POST("/worker/lookup", labor.LookupWorker)
	api.POST("/worker/:id/update", labor.UpdateWorker)
	api.POST("/worker/:id/delete", labor.DeleteWorker)
	api.POST("/worker/new", labor.AddWorker)

	// Harvest endpoints
	api.GET("/areas", harvest.AllAreas)
	api.GET("/areas/:id", harvest.GetArea)
	api.GET("/beds", harvest.AllBeds)
	api.GET("/beds/:id", harvest.GetBed)
	api.GET("/crops", harvest.AllCrops)
	api.GET("/crop/:id", harvest.GetCrop)
	api.GET("/crop/:id/preharvests", harvest.GetCropPlantings)
	api.GET("/crop/:id/harvests", harvest.GetCropHarvests)
	api.GET("/crop/:id/processing", harvest.GetCropProcessing)
	api.POST("/crop/:id/update", harvest.UpdateCrop)
	api.POST("/crop/:id/delete", harvest.DeleteCrop)
	api.POST("/crop/new", harvest.AddCrop)
	api.GET("/plantings", harvest.AllPlantings)
	api.GET("/planting/:id", harvest.GetPlanting)
	api.POST("/planting/:id/update", harvest.UpdatePlanting)
	api.POST("/planting/:id/delete", harvest.DeletePlanting)
	api.POST("/planting/new", harvest.AddPlanting)
	api.GET("/harvests", harvest.AllHarvests)
	api.GET("/harvest/:id", harvest.GetHarvest)
	api.GET("/harvest/:id/processing", harvest.GetHarvestProcessing)
	api.POST("/harvest/:id/update", harvest.UpdateHarvest)
	api.POST("/harvest/:id/delete", harvest.DeleteHarvest)
	api.POST("/harvest/new", harvest.AddHarvest)
	api.GET("/processing", harvest.AllProcessing)
	api.GET("/process/:id", harvest.GetProcessing)
	api.POST("/process/:id/update", harvest.UpdateProcessing)
	api.POST("/process/:id/delete", harvest.DeleteProcessing)
	api.POST("/process/new", harvest.AddProcessing)
	api.GET("/bin/:bin/harvest", harvest.GetBinHarvest)

	// Sales Endpoints
	api.GET("/products", sales.AllProducts)
	api.GET("/product/:id", sales.GetProduct)
	api.GET("/product/:id/sales", sales.GetProductSales)
	api.POST("/product/:id/update", sales.UpdateProduct)
	api.POST("/product/:id/delete", sales.DeleteProduct)
	api.POST("/product/new", sales.AddProduct)
	api.GET("/sales", sales.AllSales)
	api.GET("/sale/:id", sales.GetSale)
	api.POST("/sale/:id/update", sales.UpdateSale)
	api.POST("/sale/:id/delete", sales.DeleteSale)
	api.POST("/sale/new", sales.AddSale)

	err := r.Run(conf.Host + ":" + conf.Port)
	util.Check(err, "Error starting API")
}
