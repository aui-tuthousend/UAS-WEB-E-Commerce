package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	CategoryName string
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}
