package dto

type SongUpdateDTO struct {
	ID          uint   `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	ArtistID    uint   `json:"artist_id" form:"artist_id" binding:"required"`
}
