package controllers

import (
	"../repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDb) GetHistoryDriver(c *gin.Context) {
	uid := c.Param("uid")
	history, err := repositories.FindHistoryDriver(idb.DB, uid)
	//for _,h := range history{
	//	log.Print(h.CreatedAt.Date())
	//}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "get all driver history",
			"data":    history,
		})
	}
}

func (idb *InDb) GetHistoryCustomer(c *gin.Context) {
	uid := c.Param("uid")
	history, err := repositories.FindHistoryCustomer(idb.DB, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "get all driver history",
			"data":    history,
		})
	}
}
