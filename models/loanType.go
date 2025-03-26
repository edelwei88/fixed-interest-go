package models

type LoanType struct {
	ID              uint
	Type            string  `gorm:"not null"`
	Interest        float32 `gorm:"not null"`
	MinTerm         uint    `gorm:"not null"`
	MaxTerm         uint    `gorm:"not null"`
	PenaltiesPerDay float32 `gorm:"not null"`
}
