package models

import "github.com/jinzhu/gorm"

type history struct {
	gorm.Model
	OrderId    int    `json:"order_id"`
	DriverId   string `json:"driver_id"`
	CustomerId string `json:"customer_id"`
}
