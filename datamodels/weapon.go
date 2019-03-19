package datamodels

import "github.com/jinzhu/gorm"

// Weapon is a *gorm.DB type
type Weapon struct {
	gorm.Model
	Name        string `gorm:"size:100"`
	Class       string `gorm:"size:100"`
	Description string `gorm:"size:100"`
	Ability     string `gorm:"size:100"`
	Photo       string `gorm:"size:500"`
}
