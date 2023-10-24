package main

import (
	"net/http"
	"strconv"

	// Import the worker package
	"github.com/gin-gonic/gin"
	"github.com/homewardgamer/sendx-backend-IIT2020125/api"
	"github.com/homewardgamer/sendx-backend-IIT2020125/pkg/worker"
)

func main() {
	r := gin.Default()

	// Serve static files (HTML, JS, CSS)
	r.Static("/assets", "./assets")
	r.Use(api.IsPayingCustomer())
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API Routes
	r.POST("/crawl", api.HandleCrawlRequest)

	// Set worker counts for paying and non-paying workers
	r.POST("/set-workers", func(c *gin.Context) {
		paying, _ := strconv.Atoi(c.DefaultPostForm("payingWorkers", "5"))
		nonPaying, _ := strconv.Atoi(c.DefaultPostForm("nonPayingWorkers", "2"))
		worker.SetPayingWorkerCount(paying)
		worker.SetFreeWorkerCount(nonPaying)
		c.JSON(http.StatusOK, gin.H{"message": "Worker counts updated successfully"})
	})

	// Set crawl speed per worker
	r.POST("/set-crawlspeed", func(c *gin.Context) {
		rate, _ := strconv.Atoi(c.DefaultPostForm("rate", "100"))
		worker.SetRateLimit(rate)
		c.JSON(http.StatusOK, gin.H{"message": "Crawl speed updated successfully"})
	})

	// Configure routes for updating worker counts and rate limits.
	r.POST("/config/pay-workers", api.UpdatePayingWorkerCount)
	r.POST("/config/free-workers", api.UpdateFreeWorkerCount)
	r.POST("/config/rate-limit", api.UpdateRateLimit)
	r.GET("/config", api.GetCurrentConfig)

	// Start the worker processes.
	worker.StartWorkers()

	// Start the server on port 8080.
	r.Run(":8080")
}
