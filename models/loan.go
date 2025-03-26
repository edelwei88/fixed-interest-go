package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Loan struct {
	ID           uint
	InitialValue decimal.Decimal `gorm:"type:decimal(20,2);not null"`
	Time         time.Time       `gorm:"not null"`
	Term         uint            `gorm:"not null"`
	Payday       uint            `gorm:"not null;check:payday < 32"`
	LoanTypeID   uint            `gorm:"not null"`
	LoanType     LoanType
	UserID       uint `gorm:"not null"`
	LoanPayments []LoanPayment
}
