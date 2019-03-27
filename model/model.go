package model

import "github.com/jinzhu/gorm"

type DriverTest struct {
	gorm.Model
	Name       string
	Coordinate uint64
	Status     string
}
type CustomerTest struct {
	gorm.Model
	Name       string
	Coordinate uint64
	Status     string
}
