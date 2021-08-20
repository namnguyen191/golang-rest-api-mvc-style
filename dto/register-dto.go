package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}
