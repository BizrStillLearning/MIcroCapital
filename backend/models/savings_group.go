package models

import (
	"time"

	"gorm.io/gorm"
)

type SavingsGroup struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name"`
	MonthlyFee   float64        `gorm:"type:decimal(15,2);not null" json:"monthly_fee"`
	TotalPool    float64        `gorm:"type:decimal(15,2);default:0" json:"total_pool"`
	CurrentCycle int            `gorm:"default:1" json:"current_cycle"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
