package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type PhotosProduct struct {
	gorm.Model
	IdProduct  uint
	ImagePaths datatypes.JSON `json:"image_paths"`
	CreatedAt  time.Time      `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time      `gorm:"type:datetime;not null"`
}