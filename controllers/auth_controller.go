package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

func Login(c *gin.Context) {
	var request requests.LoginRequest
	c.BindJSON(&request)

	jwt, err := services.Login(c, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"token": jwt},
	})
}

func Register(c *gin.Context) {
	var request requests.RegisterRequest
	c.BindJSON(&request)

	if errs := request.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"data":    errs,
		})
		return
	}

	err := services.Register(c, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Me(c *gin.Context) {
	user, err := services.Me(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    user,
	})
}
