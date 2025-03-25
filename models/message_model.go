package models

import "time"

type Message struct {
	ID               string    `gorm:"primaryKey"`
	SenderID         string    `gorm:"not null"`
	ChatID           string    `gorm:"not null"`
	EncryptedContent string    `gorm:"not null"`
	Timestamp        time.Time `gorm:"autoCreateTime"`
	Status           string    `gorm:"type:string;default:'undelivered'"`

	Sender *User `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Chat   *Chat `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
}
