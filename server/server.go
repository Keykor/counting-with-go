package main

import (
	"counting/starter"
	"counting/worker"
	"github.com/gin-gonic/gin"
)

func createWorker(c *gin.Context) {
	worker.CreateWorker()
	c.JSON(200, gin.H{
		"message": "Worker created",
	})
}

type counterParams struct {
	Amount int `json:"amount" binding:"required,gt=0"`
}

func updateCounter(c *gin.Context) {
	var cParams counterParams

	if err := c.ShouldBindJSON(&cParams); err != nil {
		c.JSON(400, gin.H{"error": "No amount provided"})
		return
	}

	starter.UpdateCounter(cParams.Amount)
	c.JSON(200, gin.H{
		"message": "Counting started",
	})
}

func getCounter(c *gin.Context) {
	counterValue := starter.GetCounter()
	c.JSON(200, gin.H{
		"counter": counterValue,
	})
}

func resetCounter(c *gin.Context) {
	starter.ResetCounter()
	c.JSON(200, gin.H{
		"message": "Counter reset",
	})
}

func main() {
	r := gin.Default()
	r.GET("/createWorker", createWorker)
	r.POST("/updateCounter", updateCounter)
	r.GET("/getCounter", getCounter)
    r.GET("/resetCounter", resetCounter)
	r.Run()
}
