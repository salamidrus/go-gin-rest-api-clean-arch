package dto

// RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name string `json:"name" form:"name" binding:"required" validate:"min:1"`
}
