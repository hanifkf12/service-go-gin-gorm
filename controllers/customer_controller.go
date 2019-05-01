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

func (idb *InDb) GetAllCustomers(c *gin.Context) {
	customers, err := repositories.GetCustomers(idb.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    customers,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Fetch All Customers",
		"data":    customers,
	})
}

func (idb *InDb) RegisterCustomer(c *gin.Context) {
	var (
		customer models.Customer
	)
	uid := c.PostForm("uid")
	name := c.PostForm("name")
	email := c.PostForm("email")
	customer.Name = name
	customer.Uid = uid
	customer.Email = email
	cust, e := repositories.FindCustomerId(idb.DB, uid)
	token, _ := libraries.ClaimToken(uid)
	if e != nil {
		customer.CreatedAt = time.Now()
		data, err := repositories.CreateCustomer(idb.DB, customer)
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
			"data":    cust,
			"token":   token,
		})
		return
	}
}

//func (idb *InDb) UpdateCustomer(c *gin.Context) {
//
//}

func (idb *InDb) UpdateLocationCustomer(c *gin.Context) {
	var updateLoc models.Customer
	uid := c.Param("uid")
	lat, _ := strconv.ParseFloat(c.PostForm("lat"), 64)
	long, _ := strconv.ParseFloat(c.PostForm("long"), 64)
	location := c.PostForm("location")
	updateLoc.Lat = lat
	updateLoc.Long = long
	updateLoc.Location = location
	old, err := repositories.FindCustomerId(idb.DB, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "update location failed " + err.Error(),
			"data":    nil,
		})
		return
	} else {
		data, _ := repositories.UpdateCustomer(idb.DB, old, updateLoc)
		c.JSON(http.StatusOK, gin.H{
			"message": "update location succes",
			"data":    data,
		})
	}
}

func (idb *InDb) ShowCustomer(c *gin.Context) {
	uid := c.Param("uid")
	data, err := repositories.FindCustomerId(idb.DB, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "show customer failed " + err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "show customer " + uid,
		"data":    data,
	})
}

func (idb *InDb) DeleteCustomer(c *gin.Context) {
	uid := c.Param("uid")
	data, err := repositories.FindCustomerId(idb.DB, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete customer failed " + err.Error(),
		})
		return
	} else {
		_ = repositories.DeleteCustomer(idb.DB, data)
		c.JSON(http.StatusOK, gin.H{
			"message": "delete customer " + uid,
		})
		return
	}
}

func (idb *InDb) StoreTokenCustomer(c *gin.Context) {
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
