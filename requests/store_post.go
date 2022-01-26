package requests

type StorePostRequest struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Tags       []string `json:"tags"`
	PostImages []string `json:"post_images"`
}
