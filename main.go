package main

import (
	"./config"
	"./controllers"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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
	order := router.Group("/api/v1/order")
	{
		order.POST("/test", inDB.CreateOrder)
	}
	ginpprof.Wrap(router)
	err := router.Run()
	if err != nil {
		log.Print(err)
	}
}
