package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/services"
)

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
	})
}

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
			"message": "Error occured",
			"errors":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

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
			"message": "Error occured",
			"errors":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}