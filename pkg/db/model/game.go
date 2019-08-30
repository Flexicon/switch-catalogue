package model

import (
	"time"
)

type Game struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Title     string `gorm:"not null"`
}
