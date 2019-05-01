package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetCustomers(db *gorm.DB) ([]models.Customer, error) {
	var customers []models.Customer
	err := db.Find(&customers).Error
	return customers, err
}
func CreateCustomer(db *gorm.DB, customer models.Customer) (models.Customer, error) {
	err := db.Create(&customer).Error
	return customer, err
}
func DeleteCustomer(db *gorm.DB, customer models.Customer) error {
	err := db.Delete(&customer).Error
	return err
}
func ShowCustomer(db *gorm.DB, id string) (models.CustomerTest, error) {
	var customer models.CustomerTest
	err := db.First(&customer, id).Error
	return customer, err
}
func UpdateCustomer(db *gorm.DB, customer models.Customer, newCustomer models.Customer) (models.Customer, error) {
	err := db.Model(&customer).Updates(newCustomer).Error
	return customer, err
}
func FindCustomerId(db *gorm.DB, uid string) (models.Customer, error) {
	var customer models.Customer
	err := db.Where("uid = ?", uid).First(&customer).Error
	return customer, err
}
