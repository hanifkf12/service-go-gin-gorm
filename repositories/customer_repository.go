package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetCustomers(db *gorm.DB) ([]models.CustomerTest, error) {
	var customers []models.CustomerTest
	err := db.Find(&customers).Error
	return customers, err
}
func CreateCustomer(db *gorm.DB, customer models.CustomerTest) (models.CustomerTest, error) {
	err := db.Create(&customer).Error
	return customer, err
}
func DeleteCustomer(db *gorm.DB, customer models.CustomerTest) (models.CustomerTest, error) {
	err := db.Delete(&customer).Error
	return customer, err
}
func ShowCustomer(db *gorm.DB, id string) (models.CustomerTest, error) {
	var customer models.CustomerTest
	err := db.First(&customer, id).Error
	return customer, err
}
func UpdateCustomer(db *gorm.DB, customer models.CustomerTest, newCustomer models.CustomerTest) (models.CustomerTest, error) {
	err := db.Model(&customer).Updates(newCustomer).Error
	return customer, err
}
func FindCustomerId(db *gorm.DB, id string) (models.CustomerTest, error) {
	var customer models.CustomerTest
	err := db.First(&customer, id).Error
	return customer, err
}
