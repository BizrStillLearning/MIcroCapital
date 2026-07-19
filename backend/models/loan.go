package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	UserID             uint           `gorm:"not null" json:"user_id"`
	Title              string         `gorm:"type:varchar(200);not null" json:"title"`
	TotalAmount        float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	PaidAmount         float64        `gorm:"type:decimal(15,2);default:0" json:"paid_amount"`
	MonthlyInstallment float64        `gorm:"type:decimal(15,2);not null" json:"monthly_installment"`
	DueDate            time.Time      `json:"due_date"`
	Status             string         `gorm:"type:enum('pending', 'active', 'completed', 'defaulted');default:'pending'" json:"status"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID" json:"borrower,omitempty"`
}
