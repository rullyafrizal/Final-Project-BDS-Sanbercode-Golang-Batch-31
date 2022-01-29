package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

// IndexRoles godoc
// @Summary Get all roles (admin only)
// @Description Get a list of roles, only admin can do this.
// @Tags Roles
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=[]results.RoleResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/roles [get]
func IndexRoles(c *gin.Context) {
	var roles []models.Role = []models.Role{}

	if err := services.GetRoles(c, &roles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"roles": roles},
	})
}

// StoreRoles godoc
// @Summary Create new role (admin only)
// @Description Create a new role, only admin can do this.
// @Tags Roles
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param Body body models.Role true "the request body to create a new role"
// @Success 200 {object} results.JSONResult{data=results.RoleResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/roles [post]
func StoreRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)

	if err := services.CreateRole(c, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"role": role},
	})
}

// ShowRole godoc
// @Summary Get role by ID (admin only)
// @Description Create an existing role by ID, only admin can do this.
// @Tags Roles
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "the role ID"
// @Success 200 {object} results.JSONResult{data=results.RoleResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/roles/{id} [post]
func ShowRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var role models.Role

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.GetRole(c, &role, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"role": role},
	})
}

// UpdateRole godoc
// @Summary Update role by ID (admin only)
// @Description Update an existing role by ID, only admin can do this.
// @Tags Roles
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "the role ID"
// @Success 200 {object} results.JSONResult{data=results.RoleResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/roles/{id} [put]
func UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var role models.Role

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.BindJSON(&role)

	if err := services.UpdateRole(c, &role, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"role": role},
	})
}

// DestroyRole godoc
// @Summary Delete role by ID (admin only)
// @Description Delete an existing role by ID, only admin can do this.
// @Tags Roles
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "the role ID"
// @Success 200 {object} results.JSONResult{data=results.IDResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/roles/{id} [delete]
func DestroyRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.DeleteRole(c, int64(id)); err != nil {
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
