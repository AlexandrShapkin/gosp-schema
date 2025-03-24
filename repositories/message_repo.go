package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
)

type MessageRepository interface {
	Create(ctx context.Context, message *models.Message) error
	GetByID(ctx context.Context, id string) (*models.Message, error)
	GetLastInChat(ctx context.Context, chatID string) (*models.Message, error)
	// offset = 0 & limit = -1 for ignore pagination
	GetAll(ctx context.Context, offset, limit int) ([]models.Message, error)
	// offset = 0 & limit = -1 for ignore pagination
	//
	// order = ASC (for right) or DESC (for reverse) by timestamp
	GetByChat(ctx context.Context, chatID string, offset, limit int, order string) ([]models.Message, error)
	// offset = 0 & limit = -1 for ignore pagination
	//
	// order = ASC (for right) or DESC (for reverse) by timestamp
	GetFromID(ctx context.Context, chatID string, startFrom string, limit int, order string) ([]models.Message, error)
	Update(ctx context.Context, message *models.Message) error
	DeleteByID(ctx context.Context, id string) error
	DeleteByChat(ctx context.Context, chatID string) error
}
