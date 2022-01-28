package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"gorm.io/gorm"
)

func Upvote(c *gin.Context, postId int64) error {
	var vote models.Vote

	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("post_id = ? AND user_id = ? AND state = ?", postId, userId, -1).First(&vote).Error; err == nil {
		vote.State = 1
		err = db.Model(&vote).Where("post_id = ? AND user_id = ?", postId, userId).Update("state", 1).Error

		return err
	}

	if err := db.Where("post_id = ? AND user_id = ? AND state = ?", postId, userId, 1).First(&vote).Error; err == nil {
		return errors.New("user already upvoted this post")
	}

	vote.PostId = postId
	vote.UserId = userId
	vote.State = 1

	err = db.Create(&vote).Error

	return err
}

func Downvote(c *gin.Context, postId int64) error {
	var vote models.Vote

	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("post_id = ? AND user_id = ? AND state = ?", postId, userId, 1).First(&vote).Error; err == nil {
		vote.State = -1
		err = db.Model(&vote).Where("post_id = ? AND user_id = ?", postId, userId).Update("state", -1).Error

		return err
	}

	if err := db.Where("post_id = ? AND user_id = ? AND state = ?", postId, userId, -1).First(&vote).Error; err == nil {
		return errors.New("user already downvoted this post")
	}

	vote.PostId = postId
	vote.UserId = userId
	vote.State = -1

	err = db.Create(&vote).Error

	return err
}
