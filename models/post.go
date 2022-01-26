package models

import (
	"time"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

type Post struct {
	ID          int64       `json:"id" gorm:"primaryKey"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	IsPublished bool        `json:"is_published" gorm:"default:false"`
	UserId      int64       `json:"user_id"`
	PostImages  []PostImage `json:"post_images"`
	Tags        []Tag       `json:"tags" gorm:"many2many:post_tags;"`
	Reviews     []Review    `json:"reviews"`
	CreatedAt   time.Time   `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt   time.Time   `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func (p *Post) ValidatePost() map[string]string {
	var errMsg = make(map[string]string)

	for _, v := range p.PostImages {
		if !utils.IsValidUrl(v.Url) {
			errMsg["post_images"] = "invalid post image url"
		}
	}

	return errMsg
}
