package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"gorm.io/gorm"
)

func GetPosts(c *gin.Context, posts *[]models.Post) error {
	db := c.MustGet("db").(*gorm.DB)
	q := db.Preload("PostImages").Preload("Tags").Where("published_at IS NOT NULL")

	if c.Query("search") != "" {
		q = q.Where("title ILIKE ?", "%"+c.Query("search")+"%")
	}

	err := q.Find(&posts).Error

	for _, post := range *posts {
		var count int64
		db.Raw("SELECT SUM(state) FROM votes WHERE post_id = ?", post.ID).Scan(&count)

		post.VoteCount = count
	}

	return err
}

func GetMyPosts(c *gin.Context, posts *[]models.Post) error {
	db := c.MustGet("db").(*gorm.DB)
	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	q := db.Preload("PostImages").Preload("Tags").Where("user_id = ?", authId)

	if c.Query("search") != "" {
		q = q.Where("title ILIKE ?", "%"+c.Query("search")+"%")
	}

	err = q.Find(&posts).Error

	for _, post := range *posts {
		var count int64
		db.Raw("SELECT SUM(state) FROM votes WHERE post_id = ?", post.ID).Scan(&count)

		post.VoteCount = count
	}

	return err
}

func GetMyPost(c *gin.Context, post *models.Post, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	authId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	err = db.Preload("PostImages").Preload("Tags").Preload("Reviews").Preload("Reviews.User").Where("user_id = ?", authId).First(&post, id).Error

	var count int64
	db.Raw("SELECT SUM(state) FROM votes WHERE post_id = ?", id).Scan(&count)

	post.VoteCount = count

	return err
}

func GetPost(c *gin.Context, post *models.Post, id int64) error {
	db := c.MustGet("db").(*gorm.DB)
	err := db.Preload("PostImages").Preload("Tags").Preload("Reviews").Preload("Reviews.User").Where("published_at IS NOT NULL").First(&post, id).Error

	var count int64
	db.Raw("SELECT SUM(state) FROM votes WHERE post_id = ?", id).Scan(&count)

	post.VoteCount = count

	return err
}

func CreatePost(c *gin.Context, post *requests.StorePostRequest) []error {
	db := c.MustGet("db").(*gorm.DB)
	var errs []error

	db.Transaction(func(tx *gorm.DB) error {
		var tags []models.Tag
		var postImages []models.PostImage

		for _, tag := range post.Tags {
			var tagModel models.Tag = models.Tag{Name: tag}

			tx.Where(tagModel).FirstOrCreate(&tagModel)
			tags = append(tags, tagModel)
		}

		for _, url := range post.PostImages {
			var pi models.PostImage = models.PostImage{
				Url: url,
			}

			postImages = append(postImages, pi)
		}

		userId, err := utils.ExtractTokenID(c)

		if err != nil {
			return err
		}

		var post models.Post = models.Post{
			Title:      post.Title,
			Content:    post.Content,
			Tags:       tags,
			PostImages: postImages,
			UserId:     userId,
		}

		var errMsgs map[string]string = post.ValidatePost()

		if len(errMsgs) > 0 {
			for _, v := range errMsgs {
				errs = append(errs, errors.New(v))
			}

			return nil
		}

		err = tx.Save(&post).Error

		if err != nil {
			errs = append(errs, err)
		}

		return err
	})

	return errs
}

func UpdatePost(c *gin.Context, req *requests.StorePostRequest, id int64) []error {
	db := c.MustGet("db").(*gorm.DB)
	var errs []error

	db.Transaction(func(tx *gorm.DB) error {
		var post models.Post
		var tags []models.Tag
		var postImages []models.PostImage

		if err := GetMyPost(c, &post, id); err != nil {
			return err
		}

		for _, tag := range req.Tags {
			var tagModel models.Tag = models.Tag{Name: tag}

			tx.Where(tagModel).FirstOrCreate(&tagModel)
			tags = append(tags, tagModel)
		}

		post.Title = req.Title
		post.Content = req.Content

		tx.Model(&post).Association("Tags").Replace(tags)

		for _, url := range req.PostImages {
			if !utils.IsValidUrl(url) {
				errs = append(errs, errors.New("invalid post image url"))

				return errors.New("invalid post image url")
			}

			var pi models.PostImage = models.PostImage{
				Url: url,
			}

			postImages = append(postImages, pi)
		}

		tx.Model(&post).Association("PostImages").Replace(postImages)

		err := tx.Save(&post).Error

		return err
	})

	return errs
}

func DeletePost(c *gin.Context, id int64) error {
	db := c.MustGet("db").(*gorm.DB)

	db.Transaction(func(tx *gorm.DB) error {
		var post models.Post

		if err := GetMyPost(c, &post, id); err != nil {
			return err
		}

		tx.Model(&post).Association("Tags").Clear()
		tx.Model(&post).Association("PostImages").Clear()

		err := tx.Delete(&post).Error

		return err
	})

	return nil
}

func PublishPost(c *gin.Context, id int64) error {
	var post models.Post

	if err := GetMyPost(c, &post, id); err != nil {
		return err
	}

	post.PublishedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	db := c.MustGet("db").(*gorm.DB)

	err := db.Save(&post).Error

	return err
}
