package requests

type ReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
}