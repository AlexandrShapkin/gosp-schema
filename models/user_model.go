package models

import "time"

type User struct {
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex;not null"`
	DisplayName string
	AvatarUrl   string
	Status      string    `gorm:"type:string;default:'offline'"`
	LastSeen    time.Time `gorm:"autoUpdateTime"`

	Messages    []Message          `gorm:"foreignKey:SenderID"`
	ChatMembers []ChatParticipants `gorm:"foreignKey:UserID"`
}
