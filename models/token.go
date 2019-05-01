package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	Token string `json:"token"`
	Uid   string `json:"uid"`
}
