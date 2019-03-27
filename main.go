package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	db := config.DBInit()
	inDB := &controllers.InDb{DB: db}
	router := gin.Default()

	driver := router.Group("/api/v1/drivers")
	{
		driver.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   "show all available driver",
			})
		})
		driver.GET("/test", inDB.GetDrivers)
		driver.POST("/test", inDB.CreateDriver)
	}
	customer := router.Group("/api/v1/customers")
	{
		customer.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   "sho all customer",
			})
		})
	}
	router.Run()
}
