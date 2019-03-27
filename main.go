package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	driver := router.Group("/api/v1/driver")
	{
		driver.GET("/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"data":   "show all available driver",
			})
		})
	}
	router.Run()
}
