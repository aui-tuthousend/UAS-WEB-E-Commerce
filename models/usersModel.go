package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"-"`
	Saldo    float64 `gorm:"default:0"`
	Status   string
}
