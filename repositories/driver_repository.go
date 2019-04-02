package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetDriver(db *gorm.DB) []models.DriverTest {
	var drivers []models.DriverTest
	db.Find(&drivers)
	return drivers
}
