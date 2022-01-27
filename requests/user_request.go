package requests

import (
	"strings"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

type StoreUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
	Password string `json:"password" binding:"required"`
	RoleId   int64  `json:"role_id" binding:"required"`
}

func (r *StoreUserRequest) Validate() map[string]string {
	var errMsgs map[string]string = map[string]string{}

	if strings.TrimSpace(r.Email) == "" {
		errMsgs["email"] = "email is required"
	}

	if !utils.IsValidEmail(r.Email) {
		errMsgs["email"] = "invalid email"
	}

	if strings.TrimSpace(r.Avatar) == "" {
		r.Avatar = "https://ui-avatars.com/api/?background=random&name=" + r.Name
	} else {
		if !utils.IsValidUrl(r.Avatar) {
			errMsgs["avatar"] = "invalid avatar url"
		}
	}

	if strings.TrimSpace(r.Password) == "" {
		errMsgs["password"] = "password is required"
	}

	if len(r.Password) < 6 {
		errMsgs["password"] = "password must be at least 6 characters"
	}

	if strings.TrimSpace(r.Name) == "" {
		errMsgs["name"] = "name is required"
	}

	return errMsgs
}
