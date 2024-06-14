package models

import (
	"gorm.io/gorm"
	"time"
)

type PhotosProduct struct {
	gorm.Model
	IdProduct  uint
	ImagePaths string
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}
