package models

import (
	"time"
	sw "github.com/AlexandrShapkin/gosp-api-server/go"
)

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

func UserFromAPIStruct(user sw.User) User {
	t, _ := time.Parse(time.RFC3339, user.LastSeen)
	return User{
		ID: user.Id,
		Username: user.Username,
		DisplayName: user.DisplayName,
		AvatarUrl: user.AvatarUrl,
		Status: user.Status,
		LastSeen: t,
	}
}

func UserToAPIStruct(user User) sw.User {
	return sw.User{
		Id: user.ID,
		Username: user.Username,
		DisplayName: user.DisplayName,
		AvatarUrl: user.AvatarUrl,
		Status: user.Status,
		LastSeen: user.LastSeen.Format(time.RFC3339),
	}
}
