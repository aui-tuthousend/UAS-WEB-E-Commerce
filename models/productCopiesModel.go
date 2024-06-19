package models

import (
	"gorm.io/gorm"
	"time"
)

type ProductCopy struct {
	gorm.Model
	ProductName        string
	ProductDescription string
	ProductImageCover  string
	Quantity           int
	ProductPrice       int
	CreatedAt          time.Time `gorm:"type:datetime;not null"`
	UpdatedAt          time.Time `gorm:"type:datetime;not null"`
}
