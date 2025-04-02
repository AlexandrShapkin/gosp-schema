package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
)

type TokenRepository interface {
	Create(ctx context.Context, token *models.Token) error
	GetToken(ctx context.Context, tokenString string) (*models.Token, error)
	DeleteToken(ctx context.Context, tokenString string) error
	DeleteExpiredTokens(ctx context.Context) error
	DeleteTokensByUser(userID string) error
}
