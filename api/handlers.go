package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homewardgamer/sendx-backend-IIT2020125/pkg/storage"
	"github.com/homewardgamer/sendx-backend-IIT2020125/pkg/worker"
)

func HandleCrawlRequest(c *gin.Context) {
	url := c.PostForm("url")
	isPaying := c.DefaultPostForm("isPaying", "false") == "true"

	page, err := storage.GetPage(url)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "page": page})
		return
	}

	resultChan, errorChan := worker.QueueJob(url, isPaying)
	select {
	case result := <-resultChan:
		err = storage.SavePage(url, result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to save page."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "page": result})
	case err := <-errorChan:
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
	}
}

func UpdatePayingWorkerCount(c *gin.Context) {
	var requestData struct {
		Count int `json:"count" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := requestData.Count
	worker.SetPayingWorkerCount(count)
	c.JSON(http.StatusOK, gin.H{"message": "Paying worker count updated successfully."})
}

func UpdateFreeWorkerCount(c *gin.Context) {
	var requestData struct {
		Count int `json:"count" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := requestData.Count
	worker.SetFreeWorkerCount(count)
	c.JSON(http.StatusOK, gin.H{"message": "Free worker count updated successfully."})
}

func UpdateRateLimit(c *gin.Context) {
	var requestData struct {
		RateLimit int `json:"rateLimit" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rateLimit := requestData.RateLimit
	worker.SetRateLimit(rateLimit)
	c.JSON(http.StatusOK, gin.H{"message": "Rate limit updated successfully."})
}

func GetCurrentConfig(c *gin.Context) {
	config := gin.H{
		"payingWorkerCount": worker.GetPayingWorkerCount(),
		"freeWorkerCount":   worker.GetFreeWorkerCount(),
		"rateLimit":         worker.GetRateLimit(),
	}
	c.JSON(http.StatusOK, config)
}
