package models

import (
	sw "github.com/AlexandrShapkin/gosp-api-server/go"
)

type Chat struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Type string `gorm:"not null;check:type IN ('private', 'public')"`

	Messages     []Message          `gorm:"foreignKey:ChatID"`
	Participants []ChatParticipants `gorm:"foreignKey:ChatID"`
}

func ChatFromAPIStruct(chat sw.Chat) Chat {
	return Chat{
		ID: chat.Id,
		Name: chat.Name,
		Type: chat.Type,
	}
}

func ChatToAPIStruct(chat Chat) sw.Chat {
	return sw.Chat{
		Id: chat.ID,
		Name: chat.Name,
		Type: chat.Type,
	}
}