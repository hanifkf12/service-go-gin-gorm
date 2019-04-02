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
