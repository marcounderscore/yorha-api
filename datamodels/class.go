package datamodels

import "github.com/jinzhu/gorm"

// Class is a table structure.
type Class struct {
	gorm.Model
	Name string `gorm:"size:100"`
}
