package models

import (
	"gorm.io/gorm"
	"time"
)

type WishList struct {
	gorm.Model
	IdUser    uint
	Total     int
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}
