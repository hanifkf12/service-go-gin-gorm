package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CustomerTest struct {
	gorm.Model
	Name       string `json:"name"`
	Coordinate uint64 `json:"coordinate"`
	Status     string `json:"status"`
}

type Customer struct {
	Uid       string    `gorm:"primary_key",json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Telephone string    `json:"telephone"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
