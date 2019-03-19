package datamodels

import "github.com/jinzhu/gorm"

// Automata is a table structure.
type Automata struct {
	gorm.Model
	Name       string `gorm:"size:100"`
	Occupation string `gorm:"size:100"`
	Photo      string `gorm:"size:500"`
	RaceID     uint
	Race       Race
}
