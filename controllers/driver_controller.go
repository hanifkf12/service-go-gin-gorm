package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDb) GetDrivers(c *gin.Context) {
	//header := c.GetHeader("Authorization")
	//token := libraries.SplitToken(header)
	//isValid,_ := libraries.ValidateToken(token)
	//if !isValid{
	//	c.JSON(http.StatusForbidden, gin.H{
	//		"message": "unauthorized",
	//	})
	//	return
	//}
	var (
		drivers []models.Driver
		result  gin.H
		err     error
	)
	drivers, err = repositories.GetDrivers(idb.DB)
	if len(drivers) <= 0 || err != nil {
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
		driver models.Driver
		result gin.H
	)
	uid := c.PostForm("uid")
	name := c.PostForm("name")
	status := c.PostForm("status")
	lat := c.PostForm("lat")
	long := c.PostForm("long")
	//i, _ := strconv.ParseUint(coordinate, 10, 64)
	driver.Uid = uid
	driver.Name = name
	driver.Lat, _ = strconv.ParseFloat(lat, 64)
	driver.Long, _ = strconv.ParseFloat(long, 64)
	driver.Status = status
	driver.CreatedAt = time.Now()
	_, err := repositories.CreateDriver(idb.DB, driver)
	if err != nil {
		result = gin.H{
			"status": http.StatusUnprocessableEntity,
			"data":   err.Error(),
		}
	} else {
		result = gin.H{
			"status": http.StatusCreated,
			"data":   driver,
		}
	}
	c.JSON(http.StatusCreated, result)
}

func (idb *InDb) ShowDriver(c *gin.Context) {
	var (
		driver models.Driver
	)
	driver, err := repositories.ShowDriver(idb.DB, c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "driver not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": driver,
	})
}

func (idb *InDb) UpdateLocationDriver(c *gin.Context) {

	var newDriver models.Driver
	driver, err := repositories.ShowDriver(idb.DB, c.Param("uid"))
	newDriver.Lat, _ = strconv.ParseFloat(c.PostForm("lat"), 64)
	newDriver.Long, _ = strconv.ParseFloat(c.PostForm("long"), 64)
	newDriver.UpdateAt = time.Now()
	d, _ := repositories.UpdateDriver(idb.DB, driver, newDriver)
	c.JSON(http.StatusOK, gin.H{
		"driver": driver,
		"data":   d,
		"err":    err,
		"new":    newDriver,
	})
}

func (idb *InDb) UpdateStatusDriver(c *gin.Context) {
	var newDriver models.Driver
	driver, _ := repositories.ShowDriver(idb.DB, c.Param("uid"))
	newDriver.Status = c.PostForm("status")
	newDriver.UpdateAt = time.Now()
	d, _ := repositories.UpdateDriver(idb.DB, driver, newDriver)
	c.JSON(http.StatusOK, gin.H{
		"message": "status " + newDriver.Status,
		"data":    d,
	})
}
