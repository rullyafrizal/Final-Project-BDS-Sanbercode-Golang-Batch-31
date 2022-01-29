package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

// Upvote godoc
// @Summary Upvote a post
// @Description Upvote a post by inserting ID in route param, you can only do it once per post.
// @Tags Post Votes
// @Param id path string true "post id"
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.IDResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/{id}/votes/up [get]
func Upvote(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid post id",
			"data":    err.Error(),
		})
		return
	}

	err = services.Upvote(c, postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"id": postId},
	})
}

// Downvote godoc
// @Summary Downvote a post
// @Description Downvote a post by inserting ID in route param, you can only do it once per post.
// @Tags Post Votes
// @Param id path string true "post id"
// @Param Authorization header string true "Bearer token"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.IDResult}
// @Failure 500 {object} results.JSONResult{data=string}
// @Router /api/v1/posts/{id}/votes/down [get]
func Downvote(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid post id",
			"data":    err.Error(),
		})
		return
	}

	err = services.Downvote(c, postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    gin.H{"id": postId},
	})
}
