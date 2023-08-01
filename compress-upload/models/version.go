package models

import (
	"time"
)


type Version struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	IsCurrent bool      `gorm:"default:false" json:"is_current,omitempty"`
	Major     int16     `gorm:"not null" json:"major,omitempty"`
	Minor     int16     `gorm:"not null" json:"minor,omitempty"`
	Patch     string    `gorm:"not null" json:"patch,omitempty"`
	Note      string    `gorm:"size:256" json:"note,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
