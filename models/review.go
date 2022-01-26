package models

import "time"

type Review struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserId    int64     `json:"user_id"`
	PostId    int64     `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
