package controllers

import (
	"../libraries"
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"sort"
	"strconv"
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
func (idb *InDb) CreateOrderTest(c *gin.Context) {
	var order models.Order
	customerId := c.PostForm("uid")
	pickUpLat, _ := strconv.ParseFloat(c.PostForm("lat_p"), 64)
	pickUpLong, _ := strconv.ParseFloat(c.PostForm("long_p"), 64)

	log.Print("uid order:", customerId)
	log.Print("lat order :", pickUpLat)
	log.Print("long order :", pickUpLong)
	note := c.PostForm("note")
	order.CustomerId = customerId
	order.DriverId = ""
	order.PickUpLat = pickUpLat
	order.PickUpLong = pickUpLong
	order.Note = note
	availableDrivers := ShowAvailableDriver(idb.DB)
	nearDriver := GetNearDriver(availableDrivers, order)
	driverId, _ := filterDriver(nearDriver)
	fixDriver, _ := repositories.ShowDriver(idb.DB, driverId)
	c.JSON(http.StatusOK, gin.H{
		"data":   availableDrivers,
		"near":   nearDriver,
		"driver": fixDriver,
	})
	//send notif to driver

}

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
func GetNearDriver(drivers []models.Driver, order models.Order) models.Drivers {
	//var driver models.Driver
	var contents models.Drivers
	for _, d := range drivers {
		distance := libraries.CalculatDistance(d.Lat, d.Long, order.PickUpLat, order.PickUpLong)
		log.Print("distance :", distance)
		log.Print("lat 1 :", d.Lat)
		log.Print("long 1 :", d.Long)
		log.Print("lat 2 :", order.PickUpLat)
		log.Print("long 2 :", order.PickUpLong)
		if distance >= 0.0 && distance <= 3.0 {
			contents = append(contents, models.NearDriver{Uid: d.Uid, Distance: distance})
		}
	}
	return contents
}

var a = 0

func filterDriver(nearDriver models.Drivers) (string, error) {
	var uid string
	sort.Sort(nearDriver)
	if a+1 > len(nearDriver) {
		a = a - 1
	}
	uid = nearDriver[a].Uid
	a = a + 1
	return uid, errors.New("Error")
}
