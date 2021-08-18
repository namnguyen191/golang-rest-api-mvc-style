package dto

import "time"

type ArtistUpdateDTO struct {
	ID          uint      `json:"id" form:"id" binding:"required"`
	Name        string    `json:"name" form:"name" binding:"required" validate:"min:1"`
	DOB         time.Time `json:"dob" form:"dob" binding:"required"`
	Nationality string    `json:"nationality" form:"nationality" binding:"required"`
}

type ArtistCreateDTO struct {
	Name        string    `json:"name" form:"name" binding:"required" validate:"min:1"`
	DOB         time.Time `json:"dob" form:"dob" binding:"required"`
	Nationality string    `json:"nationality" form:"nationality" binding:"required" validate:"min:1"`
}
