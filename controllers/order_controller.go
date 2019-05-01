package controllers

import (
	"../libraries"
	"../models"
	"../repositories"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
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
	order.Status = "Wait"
	dataOrder, _ := repositories.CreateOrder(idb.DB, order)
	availableDrivers := ShowAvailableDriver(idb.DB)
	nearDriver := GetNearDriver(availableDrivers, order)
	sort.Sort(nearDriver)
	if len(nearDriver) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Driver Not Found In This Area",
			"data":    nearDriver,
		})
		return
	}
	go func() {
		var a = 0
		log.Println(len(nearDriver))
		for i := 0; i < len(nearDriver); i++ {
			orderN, _ := repositories.FindOrderId(idb.DB, string(dataOrder.ID))
			log.Println(orderN.Status)
			fmt.Print(orderN.Status)
			out, _ := json.Marshal(orderN)
			if orderN.Status == "Accept" {
				log.Println(orderN.Status)
				log.Println(orderN.DriverId)
				break
			}
			if a > 0 {
				if orderN.Status != "Accept" {
					PendingDriver(idb.DB, nearDriver[a-1].Uid)
				}
			}
			if i+1 > len(nearDriver) {
				i = 0
			}
			//push notif to driver
			//wait to driver accept
			//uid = nearDriver[i].Uid
			tokenFcm, _ := repositories.GetTokenByUid(idb.DB, nearDriver[i].Uid)
			da := map[string]interface{}{
				"msg":  "New Ojek Order",
				"tag":  "Order",
				"data": string(out),
			}
			om := models.Notification{
				Title: "New Ojek Order",
				Body:  "From " + orderN.CustomerId,
				Data:  da,
			}
			err := om.SendNotification(tokenFcm.Token)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 10)
			a = i
		}

		for j := 0; j < a; j++ {
			ActivateDriver(idb.DB, nearDriver[a].Uid)
		}

	}()
	c.JSON(http.StatusOK, gin.H{
		"message": "Find Driver, Please Wait...",
		"data":    nearDriver,
		"order":   dataOrder,
	})

	//driverId, _ := FilterDriver(nearDriver)
	//fixDriver, _ := repositories.FindDriverId(idb.DB, driverId)
	//order.DriverId = fixDriver.Uid
	//data, _ := repositories.CreateOrder(idb.DB, order)
	//tokenFcm,_ := repositories.GetTokenByUid(idb.DB, data.DriverId)
	//out, _ := json.Marshal(data)
	//da := map[string]interface{}{
	//	"msg": "New Ojek Order",
	//	"tag": "Order",
	//	"data": string(out),
	//}
	//om := models.Notification{
	//	Title: "New Ojek Order",
	//	Body: "From " + data.CustomerId,
	//	Data: da,
	//}
	//err := om.SendNotification(tokenFcm.Token)
	//if err!=nil{
	//	c.JSON(http.StatusOK, gin.H{
	//		"data":   data,
	//		"near":   nearDriver,
	//		"driver": fixDriver,
	//	})
	//}
	//send notif to driver
}

func (idb *InDb) ShowOrder(c *gin.Context) {
	order, err := repositories.FindOrderId(idb.DB, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    order,
			"message": "Detail Order",
		})
	}

}

//accept order
func (idb *InDb) AcceptOrder(c *gin.Context) {
	var accept models.Order
	id := c.Param("id")
	driverId := c.PostForm("driver_id")
	accept.Status = "Accept"
	accept.DriverId = driverId
	driver, _ := repositories.FindDriverId(idb.DB, driverId)
	order, _ := repositories.FindOrderId(idb.DB, id)
	if driver.Status != "Active" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "denied",
			"data":    nil,
			"driver":  driver,
			"message": "not allowed to accept order",
		})
		return
	} else {
		data, _ := repositories.UpdateOrder(idb.DB, accept, order)
		//driver, _ := repositories.FindDriverId(idb.DB, data.DriverId)
		//trigger client
		//push notification client
		c.JSON(http.StatusOK, gin.H{
			"status":  "accepted",
			"data":    data,
			"driver":  driver,
			"message": "order accepted",
		})
		return
	}

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
	var reject models.Order
	id := c.Param("id")
	reject.Status = "Reject"
	order, _ := repositories.FindOrderId(idb.DB, id)
	data, _ := repositories.UpdateOrder(idb.DB, reject, order)
	c.JSON(http.StatusOK, gin.H{
		"status":  "rejected",
		"data":    data,
		"message": "order canceled",
	})
}

func (idb *InDb) OnGoingRide(c *gin.Context) {
	var ride models.Order
	var driver models.Driver
	id := c.Param("id")
	ride.Status = "On Going"
	order, _ := repositories.FindOrderId(idb.DB, id)
	data, _ := repositories.UpdateOrder(idb.DB, ride, order)
	driver.Status = "Deliver"
	dr, _ := repositories.FindDriverId(idb.DB, data.DriverId)
	update, _ := repositories.UpdateDriver(idb.DB, dr, driver)
	c.JSON(http.StatusOK, gin.H{
		"status":  "on going ride",
		"data":    data,
		"driver":  update,
		"message": "deliver customer",
	})
}

func (idb *InDb) UpdateLocationRide(c *gin.Context) {
	id := c.Param("id")
	lat, _ := strconv.ParseFloat(c.PostForm("lat"), 64)
	long, _ := strconv.ParseFloat(c.PostForm("long"), 64)
	ride, _ := repositories.FindOrderId(idb.DB, id)
	driver, _ := repositories.FindDriverId(idb.DB, ride.DriverId)
	customer, _ := repositories.FindCustomerId(idb.DB, ride.CustomerId)
	var newD models.Driver
	var newC models.Customer
	newD.Lat = lat
	newD.Long = long
	newD.UpdatedAt = time.Now()
	newC.Lat = lat
	newC.Long = long
	newC.UpdatedAt = time.Now()
	updateD, _ := repositories.UpdateDriver(idb.DB, driver, newD)
	updateC, _ := repositories.UpdateCustomer(idb.DB, customer, newC)
	c.JSON(http.StatusOK, gin.H{
		"message": "update location ride",
		"data": gin.H{
			"driver":   updateD,
			"customer": updateC,
		},
	})
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

func ScheduleOrder(db *gorm.DB, nearDriver models.Drivers, id string) {
	//var uid string
	sort.Sort(nearDriver)
	if a+1 > len(nearDriver) {
		a = a - 1
	}
	for i := 0; i < len(nearDriver); i++ {
		order, _ := repositories.FindOrderId(db, id)
		log.Println(order.Status)
		if order.Status == "Accept" {
			log.Println(order.Status)
			break
		}
		if i+1 > len(nearDriver) {
			i = 0
		}
		//push notif to driver
		//wait to driver accept
		//uid = nearDriver[i].Uid
		time.Sleep(time.Minute * 2)
	}
}
func PendingDriver(db *gorm.DB, uid string) {
	var pending models.Driver
	pending.Status = "Pending"
	old, _ := repositories.FindDriverId(db, uid)
	updated, _ := repositories.UpdateDriver(db, old, pending)
	log.Println(updated)
}

func ActivateDriver(db *gorm.DB, uid string) {
	var pending models.Driver
	pending.Status = "Pending"
	old, _ := repositories.FindDriverId(db, uid)
	updated, _ := repositories.UpdateDriver(db, old, pending)
	log.Println(updated)
}
