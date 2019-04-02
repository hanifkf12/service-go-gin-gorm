package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (idb *InDb) CreateOrder(c *gin.Context) {
	var (
		customer models.CustomerTest
		driver   models.DriverTest
		order    models.OrderTest
	)
	idDriver := c.PostForm("id_driver")
	//idCustomer := c.PostForm("id_customer")
	//err :=idb.DB.First(&customer,idCustomer).Error
	//if err!=nil{
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "customer not found",
	//	})
	//	return
	//}
	err := idb.DB.First(&driver, idDriver).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "driver NotFound",
		})
		return
	}
	order.CustomerID = customer.ID
	order.DriverId = driver.ID
	order.DateOrder = time.Now()
	order.CoordinateCustomerLat = 2138012803981
	order.CoordinateCustomerLong = 13123123123123
	order.CoordinateDriverLat = 2309129301
	order.CoordinateDriverLong = 131231213213
	order.Note = c.PostForm("note")
	err = idb.DB.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"driver":  driver,
			"message": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"driver":  driver,
			"message": "Order Create",
			"data":    order,
		})
	}
}
