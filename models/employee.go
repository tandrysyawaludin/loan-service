package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID        uint   `gorm:"primaryKey"`
	EmpID     string `gorm:"type:varchar(255);unique;not null"`
	FullName  string `gorm:"type:varchar(255);not null"`
	Position  string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (e *Employee) TableName() string {
	return "employees"
}

// Example helper methods
func GetEmployeeByID(db *gorm.DB, empID string) (*Employee, error) {
	var employee Employee
	err := db.First(&employee, "emp_id = ?", empID).Error
	return &employee, err
}
