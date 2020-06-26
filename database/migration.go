package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// CreateSchema tables if they don't exist
func CreateSchema(db *gorm.DB) {
	log.Print("Create Schema...")
}

// Migrate the database to latest version or create if it doesn't exist
func Migrate(db *gorm.DB) {
	log.Print("Migrating DB...")

	// // Add table suffix when create tables
	// db.AutoMigrate(&User{}, &example.ExampleModel{})
	// db.Model(&User{}).AddIndex("idx_user_name", "name")

}
