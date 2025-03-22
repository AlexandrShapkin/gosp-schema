package models

import "time"

type Token struct {
	ID           string    `gorm:"primaryKey"`
	UserID       string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
