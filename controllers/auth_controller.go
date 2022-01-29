package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

// Login godoc
// @Summary Login
// @Description Login, authenticating user by email and password.
// @Tags Auth
// @Param Body body requests.LoginRequest true "the request body to authenticate user"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.TokenResult}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var request requests.LoginRequest
	c.BindJSON(&request)

	jwt, err := services.Login(c, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"token": jwt},
	})
}

// Register godoc
// @Summary Registering a new account
// @Description Register, create new user account.
// @Tags Auth
// @Param Body body requests.RegisterRequest true "the request body to create a new user account"
// @Produce json
// @Success 200 {object} results.JSONResult{data=interface{}}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/auth/register [post]
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
			"message": "Error occured",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    nil,
	})
}

// Me godoc
// @Summary Fetch authenticated user
// @Description Fetch authenticated user account data.
// @Tags Auth
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.UserResult}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/auth/me [get]
func Me(c *gin.Context) {
	user, err := services.Me(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"user": user,
		},
	})
}

// UpdatePassword godoc
// @Summary Update account password
// @Description Update user authenticated password.
// @Tags Posts
// @Param Body body requests.UpdatePasswordRequest true "the request body to update a password"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=interface{}}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/auth/update-password [put]
func UpdatePassword(c *gin.Context) {
	var request requests.UpdatePasswordRequest
	c.BindJSON(&request)

	if errs := request.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"data":    errs,
		})
		return
	}

	err := services.UpdatePassword(c, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    nil,
	})
}
