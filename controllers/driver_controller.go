package controllers

import (
	"../libraries"
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDb) GetAllDrivers(c *gin.Context) {
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

func (idb *InDb) RegisterDriver(c *gin.Context) {
	var driver models.Driver
	uid := c.PostForm("uid")
	email := c.PostForm("email")
	name := c.PostForm("name")
	photo := c.PostForm("url_photo")
	driver.Uid = uid
	driver.Email = email
	driver.Name = name
	driver.UrlPhoto = photo
	driver.Status = "Active"
	dr, e := repositories.FindDriverId(idb.DB, uid)
	token, _ := libraries.ClaimToken(uid)
	if e != nil {
		driver.CreatedAt = time.Now()
		data, err := repositories.CreateDriver(idb.DB, driver)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
				"data":    data,
				"token":   token,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login Berhasil",
				"data":    data,
				"token":   token,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Berhasil",
			"data":    dr,
			"token":   token,
		})
		return
	}
}

func (idb *InDb) ShowDriver(c *gin.Context) {
	var (
		driver models.Driver
	)
	driver, err := repositories.FindDriverId(idb.DB, c.Param("uid"))
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
	driver, err := repositories.FindDriverId(idb.DB, c.Param("uid"))
	newDriver.Lat, _ = strconv.ParseFloat(c.PostForm("lat"), 64)
	newDriver.Long, _ = strconv.ParseFloat(c.PostForm("long"), 64)
	newDriver.UpdatedAt = time.Now()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"driver":  nil,
			"message": err.Error(),
		})
		return
	} else {
		d, _ := repositories.UpdateDriver(idb.DB, driver, newDriver)
		c.JSON(http.StatusOK, gin.H{
			"driver":  d,
			"message": "updated location",
		})
		return
	}

}

func (idb *InDb) UpdateProfileDriver(c *gin.Context) {
	var newD models.Driver
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	motorNumber := c.PostForm("motor_number")
	newD.Name = name
	newD.Telephone = telephone
	newD.MotorNumber = motorNumber
	driver, err := repositories.FindDriverId(idb.DB, c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	} else {
		d, _ := repositories.UpdateDriver(idb.DB, driver, newD)
		c.JSON(http.StatusOK, gin.H{
			"message": "updated profile " + d.Name,
			"data":    d,
		})
		return
	}
}

func (idb *InDb) UpdateStatusDriver(c *gin.Context) {
	var newDriver models.Driver
	driver, err := repositories.FindDriverId(idb.DB, c.Param("uid"))
	newDriver.Status = c.PostForm("status")
	newDriver.UpdatedAt = time.Now()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	} else {
		d, _ := repositories.UpdateDriver(idb.DB, driver, newDriver)
		c.JSON(http.StatusOK, gin.H{
			"message": "update status " + d.Status,
			"data":    d,
		})
		return
	}

}
func (idb *InDb) StoreTokenDriver(c *gin.Context) {
	var store models.Token
	uid := c.PostForm("uid")
	token := c.PostForm("token")
	store.Uid = uid
	store.Token = token
	exist := repositories.CheckExistToken(idb.DB, uid)
	if exist != true {
		newt, _ := repositories.CreateToken(idb.DB, store)
		c.JSON(http.StatusCreated, gin.H{
			"message": "token stored",
			"data":    newt,
		})
	} else {
		old, _ := repositories.GetTokenByUid(idb.DB, uid)
		update, _ := repositories.UpdateToken(idb.DB, old, store)
		c.JSON(http.StatusCreated, gin.H{
			"message": "token stored",
			"data":    update,
		})
	}
}
