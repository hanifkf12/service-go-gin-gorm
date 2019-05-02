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
		driver.GET("/show/:uid", inDB.ShowDriver)
		driver.GET("/all", inDB.GetAllDrivers)
		driver.POST("/register", inDB.RegisterDriver)
		driver.PUT("/update/location/:uid", inDB.UpdateLocationDriver)
		driver.PUT("/update/status/:uid", inDB.UpdateStatusDriver)
		driver.PUT("/update/profile/:uid", inDB.UpdateProfileDriver)
		driver.GET("/history/:uid", inDB.GetHistoryDriver)
		driver.POST("/token", inDB.StoreTokenDriver)
	}
	customer := router.Group("/api/v1/customers")
	{
		customer.GET("/show/:uid", inDB.ShowCustomer)
		customer.POST("/register", inDB.RegisterCustomer)
		customer.GET("/all", inDB.GetAllCustomers)
		customer.GET("/history/:uid", inDB.GetHistoryCustomer)
		customer.PUT("/update/location/:uid", inDB.UpdateLocationCustomer)
		customer.POST("/token", inDB.StoreTokenCustomer)

	}
	order := router.Group("/api/v1/order")
	{
		order.POST("/create", inDB.CreateOrderTest)
		order.GET("/show/:id", inDB.ShowOrder)
		order.GET("/all", inDB.GetAllOrders)
		order.PUT("/accept/:id", inDB.AcceptOrder)
		order.PUT("/cancel/:id", inDB.CancelOrder)
		order.PUT("/reject/:id", inDB.RejectOrder)
		order.PUT("/ongoing/:id", inDB.OnGoingRide)
		order.PUT("/finish/:id", inDB.FinishOrder)
		order.PUT("/location/ride/:id", inDB.UpdateLocationRide)
	}
}
