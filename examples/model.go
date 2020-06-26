package example

import (
	"github.com/jinzhu/gorm"
)

// Models should only be concerned with database schema, more strict checking should be put in validator.
// More detail you can find here: http://jinzhu.me/gorm/models.html#model-definition
// HINT: If you want to split null and "", you should use *string instead of string.
// type ExampleModel struct {
// 	ID       uint       `gorm:"primary_key"`
// 	ExampleID    uint       `gorm:"primary_key"`
// 	Pin      uint       `gorm:"column:pin"`
// 	Customer string     `gorm:"column:customer"`
// 	Quantity uint       `gorm:"column:quantity;size:1024"`
// 	DueDate  *time.Time `gorm:"column:due_date"`
// 	Notes    string     `gorm:"column:notes"`
// }

// ExampleModel for the example
type ExampleModel struct {
	gorm.Model
	ExampleID    int64
	Pin      int64
	Customer string
	Quantity int64
	// DueDate  *time.time
	Notes string
}

// Migrate the job models
func Migrate(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&ExampleModel{})
}
