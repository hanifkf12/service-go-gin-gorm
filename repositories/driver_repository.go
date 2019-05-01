package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetDrivers(db *gorm.DB) ([]models.Driver, error) {
	var drivers []models.Driver
	err := db.Find(&drivers).Error
	return drivers, err
}
func CreateDriver(db *gorm.DB, driver models.Driver) (models.Driver, error) {
	err := db.Create(&driver).Error
	return driver, err
}
func UpdateDriver(db *gorm.DB, driver models.Driver, newDriver models.Driver) (models.Driver, error) {
	err := db.Model(&driver).Updates(newDriver).Error
	return driver, err
}

func DeleteDriver(db *gorm.DB, driver models.DriverTest) (models.DriverTest, error) {
	err := db.Delete(&driver).Error
	return driver, err
}
func FindDriverId(db *gorm.DB, uid string) (models.Driver, error) {
	var driver models.Driver
	err := db.Where("uid = ?", uid).First(&driver).Error
	return driver, err
}

//func ShowDriver(db *gorm.DB, id string) (models.Driver, error) {
//	var driver models.Driver
//	err := db.First(&driver, id).Error
//	return driver, err
//}
