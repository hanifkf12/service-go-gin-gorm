package controllers

import (
	"../libraries"
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"sort"
	"strconv"
)

//show all available order
func (idb *InDb) GetAllOrders(c *gin.Context) {
	orders, err := repositories.GetOrders(idb.DB)
	if err != nil || len(orders) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no data",
			"data":    orders,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "all orders",
		"data":    orders,
	})
}

//create order
func (idb *InDb) CreateOrderTest(c *gin.Context) {
	var order models.Order
	customerId := c.PostForm("uid")
	pickUpLat, _ := strconv.ParseFloat(c.PostForm("lat_p"), 64)
	pickUpLong, _ := strconv.ParseFloat(c.PostForm("long_p"), 64)
	destinationLat, _ := strconv.ParseFloat(c.PostForm("lat_d"), 64)
	destinationLong, _ := strconv.ParseFloat(c.PostForm("long_d"), 64)
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	log.Print("uid order:", customerId)
	log.Print("lat order :", pickUpLat)
	log.Print("long order :", pickUpLong)
	note := c.PostForm("note")
	order.CustomerId = customerId
	order.PickUpLat = pickUpLat
	order.PickUpLong = pickUpLong
	order.DestinationLat = destinationLat
	order.DestinationLong = destinationLong
	order.Price = price
	order.Note = note
	order.Status = "Not Accepted"
	availableDrivers := ShowAvailableDriver(idb.DB)
	nearDriver := GetNearDriver(availableDrivers, order)
	driverId, _ := FilterDriver(nearDriver)
	fixDriver, _ := repositories.ShowDriver(idb.DB, driverId)
	order.DriverId = driverId
	data, _ := repositories.CreateOrder(idb.DB, order)

	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"near":   nearDriver,
		"driver": fixDriver,
	})
	//send notif to driver
}

//accept order
func (idb *InDb) AcceptOrder(c *gin.Context) {
	var accept models.Order
	id := c.Param("id")
	accept.Status = "Accepted"
	accept.DriverId = c.PostForm("driver_id")
	order, _ := repositories.FindOrderId(idb.DB, id)
	data, _ := repositories.UpdateOrder(idb.DB, accept, order)
	driver, _ := repositories.FindDriverId(idb.DB, data.DriverId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "accepted",
		"data":    data,
		"driver":  driver,
		"message": "order accepted",
	})
}

//cancel order
func (idb *InDb) CancelOrder(c *gin.Context) {
	var cancel models.Order
	id := c.Param("id")
	cancel.Status = "Cancel"
	order, _ := repositories.FindOrderId(idb.DB, id)
	data, _ := repositories.UpdateOrder(idb.DB, cancel, order)
	c.JSON(http.StatusOK, gin.H{
		"status":  "canceled",
		"data":    data,
		"message": "order canceled",
	})
}

// reject order
func (idb *InDb) RejectOrder(c *gin.Context) {
	return
}

// finish order
func (idb *InDb) FinishOrder(c *gin.Context) {
	var finish models.Order
	id := c.Param("id")
	finish.Status = "Finish"
	order, _ := repositories.FindOrderId(idb.DB, id)
	data, _ := repositories.UpdateOrder(idb.DB, finish, order)
	c.JSON(http.StatusOK, gin.H{
		"status":  "finished",
		"data":    data,
		"message": "order finished",
	})
}

//show all available driver status
func ShowAvailableDriver(db *gorm.DB) []models.Driver {
	var (
		drivers []models.Driver
	)
	data, _ := repositories.GetDrivers(db)
	for _, driver := range data {
		if driver.Status == "Active" {
			drivers = append(drivers, driver)
		}
	}
	return drivers
}

// distance 0 - 3 KM
func GetNearDriver(drivers []models.Driver, order models.Order) models.Drivers {
	//var driver models.Driver
	var contents models.Drivers
	for _, d := range drivers {
		distance := libraries.CalculatDistance(d.Lat, d.Long, order.PickUpLat, order.PickUpLong)
		if distance >= 0.0 && distance <= 3.0 {
			contents = append(contents, models.NearDriver{Uid: d.Uid, Distance: distance})
		}
	}
	return contents
}

//sort distance driver
var a = 0

func FilterDriver(nearDriver models.Drivers) (string, error) {
	var uid string
	sort.Sort(nearDriver)
	if a+1 > len(nearDriver) {
		a = a - 1
	}
	uid = nearDriver[a].Uid
	a = a + 1
	return uid, nil
}

func (idb *InDb) HistoryDriver(c *gin.Context) {

}
