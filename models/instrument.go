package models

import "gorm.io/gorm"

// Instrument - Gorm database model
type Instrument struct {
	gorm.Model         // It has ID, CreatedAt, UpdatedAt, DeletedAt as default
	Name       string  `gorm:"not null" json:"name"`
	Trades     []Trade `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"trades"`
}
