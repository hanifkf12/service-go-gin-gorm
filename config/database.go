package config

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/shejak-test?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=she_jek_test password=hanif sslmode=disable")
	if err != nil {
		panic("failed to connect to database " + err.Error())
	}

	//db.DropTableIfExists(models.Driver{})
	db.AutoMigrate(models.DriverTest{}, models.CustomerTest{}, models.OrderTest{}, models.Driver{}, models.Order{})
	return db
}
