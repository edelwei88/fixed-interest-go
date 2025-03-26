package models

type Role struct {
	ID   uint
	Role string `gorm:"not null"`
}
