package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"gorm.io/gorm"
)

func Login(c *gin.Context, request requests.LoginRequest) (string, error) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return "", errors.New("email not found")
	}

	if !user.CheckPasswordHash(request.Password) {
		return "", errors.New("invalid password")
	}

	jwt, err := utils.GenerateToken(user.Id)

	return jwt, err
}

func Register(c *gin.Context, request requests.RegisterRequest) error {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("email = ?", request.Email).First(&user).Error; err == nil {
		return errors.New("email already exists")
	}

	var role models.Role

	db.Where("name = ?", "user").First(&role)

	user.Email = request.Email
	user.Password = request.Password
	user.RoleId = role.ID

	err := db.Create(&user).Error

	return err
}

func Me(c *gin.Context) (models.User, error) {
	var user models.User

	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		return user, err
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}
