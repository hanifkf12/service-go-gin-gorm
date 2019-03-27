package config

import (
	"github.com/jinzhu/gorm"
	"log"
)
import "../model"

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/shejak-test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
		panic("failed to connect to database")
	}

	//db.DropTableIfExists(model.DriverTest{})
	db.AutoMigrate(model.DriverTest{})
	return db
}
