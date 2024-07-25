package domain

import (
	"github.com/google/uuid"
)

// TokenPayload is an entity that represents the payload of the token
type TokenPayload struct {
	ID     uuid.UUID
	UserID uint64
	Role   UserRole
}

// AuthResponse represents an authentication response body
type AuthResponse struct {
	AccessToken string `json:"token" example:"v2.local.Gdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

// newAuthResponse is a helper function to create a response body for handling authentication data
func AewAuthResponse(token string) AuthResponse {
	return AuthResponse{
		AccessToken: token,
	}
}
