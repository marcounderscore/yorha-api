package datamodels

import "github.com/jinzhu/gorm"

// Boss is a table structure.
type Boss struct {
	gorm.Model
	Name    string `gorm:"size:100"`
	Faction string `gorm:"size:100"`
	Zones   []Zone
	Photo   string `gorm:"size:500"`
}
