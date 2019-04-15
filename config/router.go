package config

import (
	"../controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func InitRouter(router *gin.Engine, db *gorm.DB) {

	inDB := &controllers.InDb{DB: db}
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "not match url routing",
		})
	})
	driver := router.Group("/api/v1/drivers")
	{
		driver.GET("/test/:uid", inDB.ShowDriver)
		driver.GET("/test", inDB.GetDrivers)
		driver.POST("/test", inDB.CreateDriver)
		driver.PUT("/update/location/:uid", inDB.UpdateLocationDriver)
		driver.PUT("/update/status/:uid", inDB.UpdateStatusDriver)
	}
	customer := router.Group("/api/v1/customers")
	{
		customer.POST("/test", inDB.CreateCustomer)
	}
	order := router.Group("/api/v1/order")
	{
		order.POST("/test", inDB.CreateOrderTest)
	}

}
