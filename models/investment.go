package models

import "time"

type Investment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	LoanID     uint      `json:"loan_id"`
	InvestorID uint      `json:"investor_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
