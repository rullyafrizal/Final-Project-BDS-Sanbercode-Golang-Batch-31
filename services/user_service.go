package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context, users *[]models.User) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Find(users).Error

	return err
}

func GetUser(c *gin.Context, user *models.User, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.First(&user, id).Error

	return err
}

func CreateUser(c *gin.Context, user *models.User) error {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		return errors.New("email already exists")
	}

	err := db.Create(&user).Error

	return err
}

func UpdateUser(c *gin.Context, user *models.User, id int64) error {
	db := c.MustGet("db").(*gorm.DB)

	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	if authId != id {
		return errors.New("you are not authorized to update this user")
	}

	if err := db.Where("email = ?", user.Email).Where("id != ", user.Id).First(&user).Error; err == nil {
		return errors.New("email already exists")
	}

	err = db.Model(&user).Where("id = ?", id).Updates(user).Error

	return err
}

func DeleteUser(c *gin.Context, id int64) error {
	var user models.User

	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	if authId != id {
		return errors.New("you are not authorized to update this user")
	}

	if err := GetUser(c, &user, id); err != nil {
		return err
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
