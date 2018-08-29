package model

import "time"

type BasicModel struct {
	ID 			uint `gorm:"primary_key"`
	CreatedAt 	time.Time
	UpdatedAt   time.Time
}

