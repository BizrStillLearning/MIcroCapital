package models

import "time"

type Group struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	MonthlyFee float64   `gorm:"type:decimal(15,2);default:0" json:"monthly_fee"`
	TotalPool  float64   `gorm:"type:decimal(15,2);default:0" json:"total_pool"`
	CreatedAt  time.Time `json:"created_at"`
}
