package controllers

import "github.com/jinzhu/gorm"

type InDb struct {
	DB *gorm.DB
}
