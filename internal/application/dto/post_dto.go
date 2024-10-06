package dto

type CreatePost struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body"  validate:"required"`
}
