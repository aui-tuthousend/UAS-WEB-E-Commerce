package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	IdUser     uint
	TotalPrice int
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}
