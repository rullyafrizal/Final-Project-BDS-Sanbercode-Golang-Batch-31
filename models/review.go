package models

import "time"

type Review struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment,omitempty"`
	Rating    float64   `json:"rating,omitempty"`
	UserId    int64     `json:"-,omitempty"`
	PostId    int64     `json:"-,omitempty"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
