package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

// IndexUsers godoc
// @Summary Get all users (admin only)
// @Description Get a list of users, this endpoint only for admin.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=[]results.UserResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/users [get]
func IndexUsers(c *gin.Context) {
	var users []models.User = []models.User{}

	if err := services.GetUsers(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"users": users},
	})
}

// ShowUser godoc
// @Summary Get user by ID
// @Description Get an existing user byID.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID"
// @Success 200 {object} results.JSONResult{data=results.UserResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/users/{id} [get]
func ShowUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var user models.User

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.GetUser(c, &user, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"user": user},
	})
}

// StoreUser godoc
// @Summary Create new user (admin only)
// @Description Create new user, this endpoint only for admin.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body requests.StoreUserRequest true "User data"
// @Success 200 {object} results.JSONResult{data=results.UserResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/users [post]
func StoreUser(c *gin.Context) {
	var user models.User
	var request requests.StoreUserRequest
	c.BindJSON(&request)

	if errs := request.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"data":  errs,
		})
		return
	}

	user = models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Avatar:   request.Avatar,
		RoleId:   request.RoleId,
	}

	if err := services.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"user": user},
	})
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update user by ID, you can only update your own data.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body requests.UpdateUserRequest true "User data"
// @Success 200 {object} results.JSONResult{data=results.UserResult}
// @Failure 400 {object} results.JSONResult{data=[]map[string]string}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var request requests.UpdateUserRequest
	var user models.User
	c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	if errs := request.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
			"data":  errs,
		})
		return
	}

	user = models.User{
		Name:   request.Name,
		Email:  request.Email,
		Avatar: request.Avatar,
	}

	if err := services.UpdateUser(c, &user, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"user": user},
	})
}

// DestroyUser godoc
// @Summary Delete user by ID
// @Description Delete user by ID, you can only delete your own data.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID"
// @Success 200 {object} results.JSONResult{data=results.IDResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/users/{id} [delete]
func DestroyUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.DeleteUser(c, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"id": id},
	})
}
