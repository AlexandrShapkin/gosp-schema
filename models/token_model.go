package models

import (
	"time"
)

type Token struct {
	ID           string    `gorm:"primaryKey"`
	UserID       string    `gorm:"not null;index"`
	RefreshToken string    `gorm:"not null;unique"`
	ExpiresAt    time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
