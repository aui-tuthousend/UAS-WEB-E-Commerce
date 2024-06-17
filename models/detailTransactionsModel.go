package models

import (
	"gorm.io/gorm"
	"time"
)

type DetailTransaction struct {
	gorm.Model
	IdTransaction   uint
	IdProductCopies uint
	Quantity        int
	Total           int
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
}
