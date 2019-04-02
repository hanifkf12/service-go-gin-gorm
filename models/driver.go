package models

import (
	"github.com/jinzhu/gorm"
)

type DriverTest struct {
	gorm.Model
	Name       string `json:"name"`
	Coordinate uint64 `json:"coordinate"`
	Status     string `json:"status"`
}
