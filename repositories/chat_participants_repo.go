package repositories

import "context"

type ChatParticipantsRepository interface {
	AddUserToChat(ctx context.Context, chatID, userID string) error
	RemoveUserFromChat(ctx context.Context, chatID, userID string) error
}