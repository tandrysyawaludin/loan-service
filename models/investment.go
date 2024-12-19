package models

import (
	"time"

	"gorm.io/gorm"
)

type Investment struct {
	ID         uint    `gorm:"primaryKey"`
	LoanID     uint    `gorm:"not null;index"` // Foreign key to Loan
	InvestorID string  `gorm:"type:varchar(255);not null"`
	Amount     float64 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (i *Investment) TableName() string {
	return "investments"
}

// Example helper methods
func GetInvestmentsByLoanID(db *gorm.DB, loanID uint) ([]Investment, error) {
	var investments []Investment
	err := db.Where("loan_id = ?", loanID).Find(&investments).Error
	return investments, err
}

func AddInvestment(db *gorm.DB, investment *Investment) error {
	return db.Create(investment).Error
}
