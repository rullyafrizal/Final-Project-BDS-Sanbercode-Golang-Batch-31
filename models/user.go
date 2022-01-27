package models

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Avatar    string    `json:"avatar,omitempty"`
	RoleId    int64     `json:"-"`
	Posts     []Post    `json:"posts,omitempty"`
	Reviews   []Review  `json:"reviews,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = HashPassword(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) IsAdmin(c *gin.Context) bool {
	var role Role

	db := c.MustGet("db").(*gorm.DB)
	db.Where("name = ?", "admin").First(&role)

	return u.RoleId == role.ID
}

func (u *User) Validate() map[string]string {
	var errors = make(map[string]string)

	if strings.TrimSpace(u.Name) == "" {
		errors["name"] = "name can't be blank"
	}

	if strings.TrimSpace(u.Email) == "" {
		errors["email"] = "email can't be blank"
	}

	if !utils.IsValidEmail(u.Email) {
		errors["email"] = "email is invalid"
	}

	return errors
}
