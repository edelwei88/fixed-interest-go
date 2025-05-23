package models

import "time"

type Token struct {
	ID         uint
	Token      string    `gorm:"size:256;not null"`
	ExpireDate time.Time `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
}
