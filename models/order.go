package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderTest struct {
	gorm.Model
	CustomerID             uint
	DriverId               uint
	DateOrder              time.Time
	CoordinateCustomerLat  uint64
	CoordinateCustomerLong uint64
	CoordinateDriverLat    uint64
	CoordinateDriverLong   uint64
	Note                   string
}

type Order struct {
	gorm.Model
	CustomerId      string  `json:"customer_id"`
	DriverId        string  `json:"driver_id"`
	PickUpLat       float64 `json:"pick_up_lat"`
	PickUpLong      float64 `json:"pick_up_long"`
	DestinationLat  float64 `json:"destination_lat"`
	DestinationLong float64 `json:"destination_long"`
	Price           float64 `json:"price"`
	Note            string  `json:"note"`
	Status          string  `json:"status"`
}
