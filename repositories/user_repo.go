package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user *models.User) error
	UpdateStatus(ctx context.Context, id string, status string) error
	DeleteByID(ctx context.Context, id string) error
}