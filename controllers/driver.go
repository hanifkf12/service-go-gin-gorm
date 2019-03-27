package controllers

import (
	"../model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (idb *InDb) GetDrivers(c *gin.Context) {
	var (
		drivers []model.DriverTest
		result  gin.H
	)
	idb.DB.Find(&drivers)
	if len(drivers) <= 0 {
		result = gin.H{
			"status": http.StatusOK,
			"data":   drivers,
			"count":  len(drivers),
		}
	} else {
		result = gin.H{
			"status": http.StatusOK,
			"data":   drivers,
			"count":  len(drivers),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDb) CreateDriver(c *gin.Context) {
	var (
		driver model.DriverTest
		result gin.H
	)
	name := c.PostForm("name")
	coordinate := c.PostForm("coordinate")
	status := c.PostForm("status")
	i, err := strconv.ParseUint(coordinate, 10, 64)
	log.Print(err, i)
	driver.Name = name
	driver.Coordinate = i
	driver.Status = status
	idb.DB.Create(&driver)
	result = gin.H{
		"status": http.StatusCreated,
		"data":   driver,
	}
	c.JSON(http.StatusCreated, result)
}
