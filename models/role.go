package models

type Role struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name,omitempty"`
	User []User `json:"users,omitempty"`
}
