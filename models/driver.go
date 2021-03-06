package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DriverTest struct {
	gorm.Model
	Name       string `json:"name"`
	Coordinate uint64 `json:"coordinate"`
	Status     string `json:"status"`
}

type Driver struct {
	Uid         string    `gorm:"primary_key" json:"uid"`
	Email       string    `gorm:"not_null" json:"email"`
	Name        string    `gorm:"not_null" json:"name"`
	UrlPhoto    string    `json:"url_photo"`
	Status      string    `json:"status"`
	Telephone   string    `json:"telephone"`
	MotorNumber string    `json:"motor_number"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NearDriver struct {
	Uid      string  `json:"uid"`
	Distance float64 `json:"distance"`
}

type Drivers []NearDriver

func (d Drivers) Len() int {
	return len(d)
}

func (d Drivers) Less(i, j int) bool {
	return d[i].Distance < d[j].Distance
}

func (d Drivers) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
