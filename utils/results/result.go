package results

import (
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/requests"
)

type JSONResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IDResult struct {
	ID int64 `json:"id"`
}

type TokenResult struct {
	Token string `json:"token"`
}

type UserResult struct {
	User models.User `json:"user"`
}

type ReviewResult struct {
	Review requests.ReviewRequest `json:"review"`
}

type RoleResult struct {
	Role models.Role `json:"role"`
}