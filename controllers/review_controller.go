package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

// StoreReview godoc
// @Summary Create new post review
// @Description Create a new post review.
// @Tags Post Reviews
// @Param Body body requests.ReviewRequest true "the request body to create a new post review"
// @Param Authorization header string true "Bearer token"
// @Param id path string true "post id"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.ReviewResult}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/{id}/reviews [post]
func StoreReview(c *gin.Context) {
	var request requests.ReviewRequest
	c.BindJSON(&request)

	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data": err.Error(),
		})
		return
	}

	if err := services.CreateReview(c, postId, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"errors":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"review": request,
		},
	})
}

// StoreReview godoc
// @Summary Update post review
// @Description Update an existing post review by ID.
// @Tags Post Reviews
// @Param Body body requests.ReviewRequest true "the request body to update a post review"
// @Param Authorization header string true "Bearer token"
// @Param id path string true "post id"
// @Param review_id path string true "review id"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.ReviewResult}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/{id}/reviews/{review_id} [put]
func UpdateReview(c *gin.Context) {
	var request requests.ReviewRequest
	c.BindJSON(&request)

	reviewId, err := strconv.ParseInt(c.Param("review_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data": err.Error(),
		})
		return
	}

	if err := services.UpdateReview(c, reviewId, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"review": request,
		},
	})
}

// DestroyReview godoc
// @Summary Delete post review
// @Description Delete an existing post review by ID.
// @Tags Post Reviews
// @Param Authorization header string true "Bearer token"
// @Param id path string true "post id"
// @Param review_id path string true "review id"
// @Produce json
// @Success 200 {object} results.JSONResult{data=results.IDResult}
// @Failure 500 {object} results.JSONResult{}
// @Router /api/v1/posts/{id}/reviews/{review_id} [delete]
func DestroyReview(c *gin.Context) {
	reviewId, err := strconv.ParseInt(c.Param("review_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured",
			"data": err.Error(),
		})
		return
	}

	if err := services.DeleteReview(c, reviewId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
			"message": "Error occured",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": gin.H{
			"id": reviewId,
		},
	})
}