package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"gorm.io/gorm"
)

func GetRoles(c *gin.Context, roles *[]models.Role) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Find(roles).Error

	return err
}

func GetRole(c *gin.Context, role *models.Role, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.First(&role, id).Error

	return err
}

func CreateRole(c *gin.Context, role *models.Role) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Create(&role).Error

	return err
}

func UpdateRole(c *gin.Context, role *models.Role, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Model(&role).Where("id = ?", id).Updates(role).Error

	return err
}

func DeleteRole(c *gin.Context, id int64) error {
	var role models.Role

	if err := GetRole(c, &role, id); err != nil {
		return err
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}

	return nil
}
