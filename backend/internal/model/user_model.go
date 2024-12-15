package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(32);not null;unique"`
	Password  string `gorm:"type:varchar(128);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
