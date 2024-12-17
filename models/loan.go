package models

import "time"

type Loan struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	BorrowerID      uint      `json:"borrower_id"`
	PrincipalAmount float64   `json:"principal_amount"`
	State           string    `json:"state"`
	Rate            float64   `json:"rate"`
	ROI             float64   `json:"roi"`
	AgreementLink   string    `json:"agreement_link"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
