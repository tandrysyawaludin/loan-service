package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID            uint    `gorm:"primaryKey"`
	BorrowerID    string  `gorm:"type:varchar(255);not null"`
	Principal     float64 `gorm:"not null"`
	Rate          float64 `gorm:"not null"`
	ROI           float64 `gorm:"not null"`
	State         string  `gorm:"type:varchar(50);not null"` // e.g., "proposed", "approved"
	AgreementLink string  `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (l *Loan) TableName() string {
	return "loans"
}

// Example helper methods
func GetLoanByID(db *gorm.DB, id string) (*Loan, error) {
	var loan Loan
	err := db.First(&loan, "id = ?", id).Error
	return &loan, err
}

func UpdateLoan(db *gorm.DB, loan *Loan) error {
	return db.Save(loan).Error
}
