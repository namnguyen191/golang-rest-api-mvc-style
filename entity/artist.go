package entity

import (
	"time"
)

type Artist struct {
	ID          uint      `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	DOB         time.Time `json:"dob"`
	Nationality string    `json:"nationality"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
