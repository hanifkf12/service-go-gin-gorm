package repositories

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetToken(db *gorm.DB) ([]models.Token, error) {
	var tokens []models.Token
	err := db.Find(&tokens).Error
	return tokens, err
}

func GetTokenByUid(db *gorm.DB, uid string) (models.Token, error) {
	var token models.Token
	err := db.Where("uid = ?", uid).First(&token).Error
	return token, err
}
func UpdateToken(db *gorm.DB, token models.Token, new models.Token) (models.Token, error) {
	err := db.Model(&token).Updates(new).Error
	return token, err
}

func CreateToken(db *gorm.DB, token models.Token) (models.Token, error) {
	err := db.Create(&token).Error
	return token, err
}

func CheckExistToken(db *gorm.DB, uid string) bool {
	var token models.Token
	err := db.Where("uid = ?", uid).First(&token).Error
	if err != nil {
		return false
	}
	return true
}
