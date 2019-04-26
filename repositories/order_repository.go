package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func CreateOrder(db *gorm.DB, order models.Order) (models.Order, error) {
	err := db.Create(&order).Error
	return order, err
}
func GetOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Find(&orders).Error
	return orders, err
}
func FindOrderId(db *gorm.DB, id string) (models.Order, error) {
	var order models.Order
	err := db.First(&order, id).Error
	return order, err
}
func UpdateOrder(db *gorm.DB, new models.Order, order models.Order) (models.Order, error) {
	err := db.Model(&order).Updates(new).Error
	return order, err
}
func FindHistoryDriver(db *gorm.DB, uid string) ([]models.Order, error) {
	var history []models.Order
	err := db.Where("driver_id = ? AND status = ? ", uid, "Finish").Find(&history).Error
	return history, err
}
func FindHistoryCustomer(db *gorm.DB, uid string) ([]models.Order, error) {
	var history []models.Order
	err := db.Where("customer_id = ? AND status = ?", uid, "Finish").Find(&history).Error
	return history, err
}
