package migrator

import (
	"fmt"

	"github.com/sahlannasution/xnews-xapiens-backend/models"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {

	// Drop Table Users
	db.Migrator().DropTable(&models.Users{})

	if check := db.Migrator().HasTable(&models.Users{}); !check {
		db.Migrator().CreateTable(&models.Users{})
		fmt.Println("Table User has been created!")
	}
}
