package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (idb *InDb) CreateCustomer(c *gin.Context) {
	var (
		customer models.CustomerTest
	)
	customer.Name = c.PostForm("name")
	i, _ := strconv.ParseUint(c.PostForm("coordinate"), 10, 64)
	customer.Coordinate = i
	customer.Status = c.PostForm("status")
	err := idb.DB.Create(&customer).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err,
		})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "create",
			"data":    customer,
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
