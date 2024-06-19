package models

import "gorm.io/gorm"

type PhotoProduct struct {
	gorm.Model
	IdProduct uint
	ImgPath   string
}
