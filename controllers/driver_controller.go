package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDb) GetDrivers(c *gin.Context) {
	var (
		drivers []models.DriverTest
		result  gin.H
	)
	drivers = repositories.GetDriver(idb.DB)
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
		driver models.DriverTest
		result gin.H
	)
	name := c.PostForm("name")
	coordinate := c.PostForm("coordinate")
	status := c.PostForm("status")
	i, _ := strconv.ParseUint(coordinate, 10, 64)
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
