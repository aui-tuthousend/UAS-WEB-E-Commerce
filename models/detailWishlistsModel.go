package models

import (
	"gorm.io/gorm"
	"time"
)

type DetailWishlist struct {
	gorm.Model
	IdWishlist uint
	IdProduct  uint
	Quantity   int
	Total      int
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}
