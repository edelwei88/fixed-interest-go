package models

type User struct {
	ID           uint
	FirstName    string `gorm:"not null;size:50"`
	LastName     string `gorm:"not null;size:50"`
	PhoneNumber  string `gorm:"not null;size:10"`
	Login        string `gorm:"not null;size:256"`
	PasswordHash string `gorm:"not null;size:64"`
	RoleID       uint   `gorm:"not null"`
	Role         Role
	Loans        []Loan
}
