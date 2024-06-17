package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ProductName        string
	ProductDescription string
	ProductImageCover  string
	ProductStock       int
	ProductPrice       int
	ImagePaths         datatypes.JSON `json:"image_paths"`
	CreatedAt          time.Time      `gorm:"type:datetime;not null"`
	UpdatedAt          time.Time      `gorm:"type:datetime;not null"`
}
