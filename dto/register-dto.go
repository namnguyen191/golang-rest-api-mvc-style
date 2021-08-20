package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required,min=2"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min=6"`
}
