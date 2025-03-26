package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type LoanPayment struct {
	ID     uint
	Amount decimal.Decimal `gorm:"type:decimal(20,2);not null"`
	Time   time.Time       `gorm:"not null"`
	LoanID uint            `gorm:"not null"`
}
