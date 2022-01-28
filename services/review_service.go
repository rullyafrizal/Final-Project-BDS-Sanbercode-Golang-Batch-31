package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"gorm.io/gorm"
)

func CreateReview(c *gin.Context, postId int64, request requests.ReviewRequest) error {
	db := c.MustGet("db").(*gorm.DB)
	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	review := models.Review{
		UserId:  authId,
		PostId:  postId,
		Comment: request.Comment,
	}

	if err := db.Create(&review).Error; err != nil {
		return err
	}

	return nil
}

func UpdateReview(c *gin.Context, id int64, request requests.ReviewRequest) error {
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	db.First(&review, id)

	if review.UserId != authId {
		return errors.New("you are not authorized to update this review")
	}

	review.Comment = request.Comment

	if err := db.Save(&review).Error; err != nil {
		return err
	}

	return nil
}

func DeleteReview(c *gin.Context, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	db.First(&review, id)

	if review.UserId != authId {
		return errors.New("you are not authorized to delete this review")
	}

	if err := db.Delete(&review).Error; err != nil {
		return err
	}

	return nil
}
