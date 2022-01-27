package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

func IndexUsers(c *gin.Context) {
	var users []models.User = []models.User{}

	if err := services.GetUsers(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"users": users},
	})
}

func ShowUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var user models.User

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := services.GetUser(c, &user, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"user": user},
	})
}

func StoreUser(c *gin.Context) {
	var user models.User
	var request requests.StoreUserRequest
	c.BindJSON(&request)

	if errs := request.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"errors":  errs,
		})
		return
	}

	user = models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Avatar:  request.Avatar,
		RoleId:  request.RoleId,
	}

	if err := services.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var user models.User
	c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if errs := user.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"errors":  errs,
		})
		return
	}

	if err := services.UpdateUser(c, &user, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DestroyUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := services.DeleteUser(c, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
