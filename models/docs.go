package models

type Docs struct {
	UserID uint `gorm:"not null"`
	User   User
	Data   string `gorm:"type:bytea;not null"`
}
