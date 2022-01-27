package requests

type ReviewRequest struct {
	Comment string `json:"comment" binding:"required"`
	Rating  float64    `json:"rating" binding:"required"`
}