package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"web_uas/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("failed to connect")
	} else {
		fmt.Print("SUCK SEED!")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.PhotosProduct{})
}

func GetDB() *gorm.DB {
	return DB
}
