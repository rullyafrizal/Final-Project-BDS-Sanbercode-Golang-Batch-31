package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

// IndexPosts godoc
// @Summary Get all published posts
// @Description Get a list of published posts.
// @Tags Posts
// @Produce json
// @Success 200 {object} results.JSONResult{data=[]models.Post}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts [get]
func IndexPosts(c *gin.Context) {
	var posts []models.Post

	if err := services.GetPosts(c, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"posts": posts},
	})
}

// ShowPost godoc
// @Summary Get published post by ID
// @Description Get a post by inserting ID in route param (only published post will return the result).
// @Tags Posts
// @Param id path string true "post id"
// @Produce json
// @Success 200 {object} results.JSONResult{data=models.Post}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/{id} [get]
func ShowPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.GetPost(c, &post, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"post": post},
	})
}

// StorePost godoc
// @Summary Create new post
// @Description Create a new post.
// @Tags Posts
// @Param Body body requests.StorePostRequest true "the request body to create a new post"
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} results.JSONResult{data=requests.StorePostRequest}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/ [post]
func StorePost(c *gin.Context) {
	var request requests.StorePostRequest
	c.BindJSON(&request)

	if errs := services.CreatePost(c, &request); len(errs) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data":    utils.StringifyErrors(errs...),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"post": request,
		},
	})
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update an existing post by ID, you can only update your own post.
// @Tags Posts
// @Param id path string true "post id"
// @Param Body body requests.StorePostRequest true "the request body to update a post"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=requests.StorePostRequest}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/{id} [put]
func UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var request requests.StorePostRequest
	c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	if errs := services.UpdatePost(c, &request, int64(id)); len(errs) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data":    utils.StringifyErrors(errs...),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"post": request,
		},
	})
}

// DestroyPost godoc
// @Summary Delete a post
// @Description Delete an existing post by ID, you can only delete your own post.
// @Tags Posts
// @Produce json
// @Param id path string true "post id"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=string}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/{id} [delete]
func DestroyPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.DeletePost(c, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"id": id,
		},
	})
}

// IndexMyPosts godoc
// @Summary Get all authenticated user posts
// @Description Get a list of authenticated user posts.
// @Tags Posts
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=[]models.Post}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/my [get]
func IndexMyPosts(c *gin.Context) {
	var posts []models.Post

	if err := services.GetMyPosts(c, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"posts": posts},
	})
}

// ShowPost godoc
// @Summary Get authenticated user post by ID
// @Description Get user authenticated post by inserting ID in route param.
// @Tags Posts
// @Param id path string true "post id"
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} results.JSONResult{data=models.Post}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/{id}/my [get]
func ShowMyPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var post models.Post

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.GetMyPost(c, &post, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"post": post},
	})
}

// UpdatePost godoc
// @Summary Publish a post
// @Description Publish an existing post by ID, you can only publish your own post.
// @Tags Posts
// @Param id path string true "post id"
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} results.JSONResult{data=int}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/{id}/publish [patch]
func PublishPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	if err := services.PublishPost(c, int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"id": id,
		},
	})
}
