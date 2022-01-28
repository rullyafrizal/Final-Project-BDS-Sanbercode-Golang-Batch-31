package requests

import (
	"strings"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/utils"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterRequest) Validate() map[string]string {
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

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (u *UpdatePasswordRequest) Validate() map[string]string {
	var errMsgs map[string]string = map[string]string{}

	if strings.TrimSpace(u.OldPassword) == "" {
		errMsgs["old_password"] = "old password is required"
	}

	if strings.TrimSpace(u.NewPassword) == "" {
		errMsgs["new_password"] = "new password is required"
	}

	if len(u.NewPassword) < 6 {
		errMsgs["new_password"] = "new password must be at least 6 characters"
	}

	return errMsgs
}
