package controllers

import (
	"../libraries"
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"net/http"
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
	customer.CreatedAt = time.Now()
	cust, e := repositories.FindCustomerId(idb.DB, uid)
	token, _ := libraries.ClaimToken(uid)
	if e != nil {
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

func (idb *InDb) UpdateCustomer(c *gin.Context) {

}

func (idb *InDb) UpdateLocationCustomer(c *gin.Context) {

}

func (idb *InDb) ShowCustomer(c *gin.Context) {

}

func (idb *InDb) DeleteCustomer(c *gin.Context) {

}
