package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
