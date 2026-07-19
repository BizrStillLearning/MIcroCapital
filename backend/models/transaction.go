package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Type        string         `gorm:"type:enum('topup', 'withdraw', 'fund_campaign');not null" json:"type"`
	Amount      float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	ReferenceID uint           `json:"reference_id"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
