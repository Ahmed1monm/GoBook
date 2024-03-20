package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
