package models

import (
	"gorm.io/gorm"
	"time"
)

type DetailWishlist struct {
	gorm.Model
	IdWishlist   uint
	IdProduct    uint
	ProductName  string
	ProductPrice int
	Quantity     int       `gorm:"default: 1"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}
