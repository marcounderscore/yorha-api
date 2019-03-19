package datamodels

import "github.com/jinzhu/gorm"

// PodProgram is a table structure.
type PodProgram struct {
	gorm.Model
	Name     string `gorm:"size:100"`
	Program  string `gorm:"size:100"`
	Cooldown int
	Photo    string `gorm:"size:500"`
}
