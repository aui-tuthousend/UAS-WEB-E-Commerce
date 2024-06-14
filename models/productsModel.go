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
	IdSeller           int
	ProductStock       int
	ProductPrice       int
	CreatedAt          time.Time `gorm:"type:datetime;not null"`
	UpdatedAt          time.Time `gorm:"type:datetime;not null"`
}
