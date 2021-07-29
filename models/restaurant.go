package models

import (
	"time"
)

type Restaurant struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Type      string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
