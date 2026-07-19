package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(100);not null" json:"name"`
	Phone      string         `gorm:"type:varchar(20);unique;not null" json:"phone"`
	PIN        string         `gorm:"type:varchar(255);not null" json:"-"`
	Email      string         `gorm:"uniqueIndex;default:null" json:"email"`
	Password   string         `json:"-"`
	Role       string         `gorm:"type:enum('member', 'agent', 'admin');default:'member'" json:"role"`
	Balance    float64        `gorm:"type:decimal(15,2);default:0" json:"balance"`
	IsVerified bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
