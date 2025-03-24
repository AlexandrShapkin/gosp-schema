package gospschema

import "github.com/AlexandrShapkin/gosp-schema/repositories"

type Repository interface {
	Chat() repositories.ChatRepository
	ChatParticipants() repositories.ChatParticipantsRepository
	Message() repositories.MessageRepository
	Token() repositories.TokenRepository
	User() repositories.UserRepository
}