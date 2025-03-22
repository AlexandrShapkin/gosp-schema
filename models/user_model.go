package models

import "time"

type User struct {
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex;not null"`
	DisplayName string
	AvatarUrl   string
	Status      string    `gorm:"type:enum('online','offline');default:'offline'"`
	LastSeen    time.Time `gorm:"autoUpdateTime"`

	Messages    []Message         `gorm:"foreignKey:SenderID"`
	ChatMembers []ChatPaticipants `gorm:"foreignKey:UserID"`
}
