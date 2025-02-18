package port

import (
	"github.com/sugaml/authserver/internal/core/domain"
)

//go:generate mockgen -source=auth.go -destination=mock/auth.go -package=mock

// TokenService is an interface for interacting with token-related business logic
type TokenService interface {
	// CreateToken creates a new token for a given user
	CreateToken(uid uint) (string, error)
	// VerifyToken verifies the token and returns the payload
	VerifyToken(token string) (*domain.TokenPayload, error)
}
