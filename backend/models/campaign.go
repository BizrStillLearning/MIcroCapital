package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	Title         string         `gorm:"type:varchar(200);not null" json:"title"`
	Description   string         `gorm:"type:text;not null" json:"description"`
	TargetAmount  float64        `gorm:"type:decimal(15,2);not null" json:"target_amount"`
	CurrentAmount float64        `gorm:"type:decimal(15,2);default:0" json:"current_amount"`
	DurationDays  int            `gorm:"not null" json:"duration_days"`
	Status        string         `gorm:"type:enum('pending', 'active', 'funded', 'completed');default:'pending'" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID" json:"raiser,omitempty"`
}
