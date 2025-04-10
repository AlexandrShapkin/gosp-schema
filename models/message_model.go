package models

import (
	sw "github.com/AlexandrShapkin/gosp-api-server/go"
	"time"
)

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

func MessageFromAPIStruct(message sw.Message) Message {
	t, _ := time.Parse(time.RFC3339, message.Timestamp)
	return Message{
		ID:               message.Id,
		SenderID:         message.SenderId,
		ChatID:           message.ChatId,
		EncryptedContent: message.EncryptedContent,
		Timestamp:        t,
		Status:           message.Status,
	}
}

func MessageToAPIStatus(message Message) sw.Message {
	return sw.Message{
		Id:               message.ID,
		SenderId:         message.SenderID,
		ChatId:           message.ChatID,
		EncryptedContent: message.EncryptedContent,
		Timestamp:        message.Timestamp.Format(time.RFC3339),
		Status:           message.Status,
	}
}
