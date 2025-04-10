package models

import (
	sw "github.com/AlexandrShapkin/gosp-api-server/go"
)

type ChatParticipants struct {
	ChatID string `gorm:"primaryKey"`
	UserID string `gorm:"primaryKey"`

	Chat *Chat `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE"`
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func ChatParticipantsFromAPIStruct(chat sw.Chat, user sw.User) ChatParticipants {
	return ChatParticipants{
		ChatID: chat.Id,
		UserID: user.Id,
	}
}

func ChatParticipantsToAPIStruct(chatParticipants ChatParticipants) (sw.Chat, sw.User) {
	return ChatToAPIStruct(*chatParticipants.Chat), UserToAPIStruct(*chatParticipants.User)
}
