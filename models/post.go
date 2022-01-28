package models

import (
	"database/sql"
	"time"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

type Post struct {
	ID          int64        `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	VoteCount   int64        `json:"vote_count"`
	UserId      int64        `json:"user_id"`
	PostImages  []PostImage  `json:"post_images,omitempty"`
	Tags        []Tag        `json:"tags,omitempty" gorm:"many2many:post_tags;"`
	Reviews     []Review     `json:"reviews,omitempty"`
	Votes       []User       `json:"votes,omitempty" gorm:"many2many:votes;"`
	PublishedAt sql.NullTime `json:"published_at"`
	CreatedAt   time.Time    `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt   time.Time    `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
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
