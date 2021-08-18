package entity

type Song struct {
	ID          uint   `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	ArtistID    uint   `gorm:"not null" json:"_"`
	Artist      Artist `gorm:"foreignkey:ArtistID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"artist"`
}
