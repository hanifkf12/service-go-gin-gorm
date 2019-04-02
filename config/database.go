package config

import (
	"../models"
	"github.com/jinzhu/gorm"
	"log"
)

func DBInit() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/shejak-test?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=she_jek_test password=hanif sslmode=disable")
	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}

	//db.DropTableIfExists(models.OrderTest{})
	db.AutoMigrate(models.DriverTest{}, models.CustomerTest{}, models.OrderTest{})
	return db
}
