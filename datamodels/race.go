package datamodels

import "github.com/jinzhu/gorm"

// Race is a table structure.
type Race struct {
	gorm.Model
	Name string `gorm:"size:100"`
}
