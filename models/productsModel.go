package models

import (
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
	CategoryId         uint
	Sold               int       `gorm:"default:0"`
	CreatedAt          time.Time `gorm:"type:datetime;not null"`
	UpdatedAt          time.Time `gorm:"type:datetime;not null"`
}
