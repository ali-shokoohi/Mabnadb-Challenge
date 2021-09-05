package models

import "gorm.io/gorm"

type Instrument struct {
	gorm.Model        // It has ID, CreatedAt, UpdatedAt, DeletedAt as default
	Name       string `gorm:"not null" json:"name"`
}
