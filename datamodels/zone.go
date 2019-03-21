package datamodels

import "github.com/jinzhu/gorm"

// Zone is a table structure.
type Zone struct {
	gorm.Model
	Name string `gorm:"size:100"`
}
