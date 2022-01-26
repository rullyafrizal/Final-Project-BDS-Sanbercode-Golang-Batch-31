package models

type PostImage struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Url    string `json:"url"`
	PostId int64  `json:"post_id"`
}
