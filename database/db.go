package database

import (
	"github.com/brennschlus/house_points_server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {

	_db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = _db
	if err := db.AutoMigrate(&models.Faculty{}); err != nil {
		panic(err)
	}

	db = _db
	if err := db.AutoMigrate(&models.Branch{}); err != nil {
		panic(err)
	}

}

func GetDB() *gorm.DB {
	return db
}
