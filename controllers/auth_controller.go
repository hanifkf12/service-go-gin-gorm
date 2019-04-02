package controllers

import (
	"../libraries"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (idb *InDb) GenerateToken(c gin.Context) {
	var (
		email, username, role string
	)
	email = c.PostForm("email")
	username = c.PostForm("username")
	role = c.PostForm("role")
	//err := idb.DB.Where("id = ?", id).First(&person).Error
	token, _ := libraries.ClaimToken(email, username, role)
	fmt.Print(token)
}
