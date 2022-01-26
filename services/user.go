package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
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
	var role models.Role

	db := c.MustGet("db").(*gorm.DB)

	db.Where("name = ?", "user").First(&role)

	user.RoleId = role.ID

	err := db.Create(&user).Error

	return err
}

func UpdateUser(c *gin.Context, user *models.User, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Model(&user).Where("id = ?", id).Updates(user).Error

	return err
}

func DeleteUser(c *gin.Context, id int64) error {
	var user models.User

	if err := GetUser(c, &user, id); err != nil {
		return err
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
