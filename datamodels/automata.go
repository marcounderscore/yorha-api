package datamodels

import "github.com/jinzhu/gorm"

// Automata is a *gorm.DB type
type Automata struct {
	gorm.Model
	Name       string `gorm:"size:100"`
	Occupation string `gorm:"size:100"`
	Race       string `gorm:"size:100"`
}
