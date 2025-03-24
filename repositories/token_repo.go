package repositories

import (
	"context"

	"github.com/AlexandrShapkin/gosp-schema/models"
)

type TokenRepository interface {
	Create(ctx context.Context, token *models.Token) error
	GetToken(ctx context.Context, tokenString string) (*models.Token, error)
	DeleteToken(tokenString string) error
	DeleteExpiredTokens() error
	DeleteTokensByUser(userID string) error
}
