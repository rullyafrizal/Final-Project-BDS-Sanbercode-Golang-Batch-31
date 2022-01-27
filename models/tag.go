package models

type Tag struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name,omitempty"`
}