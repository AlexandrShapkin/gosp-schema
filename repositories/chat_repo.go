package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *models.Chat) error
	GetByID(ctx context.Context, id string) (*models.Chat, error)
	// offset = 0 & limit = -1 for ignore pagination
	GetAll(ctx context.Context, offset, limit int) ([]models.Chat, error)
	// offset = 0 & limit = -1 for ignore pagination
	GetByUser(ctx context.Context, userID string, offset, limit int) ([]models.Chat, error)
	Update(ctx context.Context, chat *models.Chat) error
	UpdateName(ctx context.Context, id string, name string) error
	DeleteByID(ctx context.Context, id string) error
}