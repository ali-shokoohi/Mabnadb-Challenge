package models

import (
	"time"

	"gorm.io/gorm"
)

// Trade - Gorm database model
type Trade struct {
	gorm.Model             // It has ID, CreatedAt, UpdatedAt, DeletedAt as default
	Open         float32   `gorm:"not null" json:"open"`
	High         float32   `gorm:"not null" json:"high"`
	Low          float32   `gorm:"not null" json:"low"`
	Close        float32   `gorm:"not null" json:"close"`
	DateEn       time.Time `gorm:"not null" json:"date_en"`
	InstrumentID uint      `gorm:"not null" json:"instrument_id"`
}
