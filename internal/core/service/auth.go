package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

type ClientStotre struct {
	repo repository.IRepository
}

// JWTClaims defines the structure of JWT token claims
type JWTClaims struct {
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	jwt.StandardClaims
}

var jwtSecret = []byte("supersecretkey")

func newClientStotreService(repo repository.IRepository) *ClientStotre {
	return &ClientStotre{
		repo: repo,
	}
}

// GetByID retrieves a client by its ID from the database
func (s *ClientStotre) GetByID(id string) (oauth2.ClientInfo, error) {
	domain := "http://localhost:9094"
	secret := ""
	result, err := s.repo.Client().GetCliendID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	results, _, err := s.repo.ClientSecret().ListByClientID(context.Background(), result.ID)
	if err != nil {
		return nil, err
	}
	for _, res := range results {
		secret = res.Value
	}
	return &models.Client{
		ID:     result.ID,
		Secret: secret,
		Domain: domain,
	}, nil
}

type JWTAccessGenerate struct{}

// Token implements the AccessGenerate interface
func (g *JWTAccessGenerate) Token(data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	now := time.Now()
	claims := JWTClaims{
		ClientID: data.Client.GetID(),
		Scope:    data.TokenInfo.GetScope(),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err = token.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token (optional)
	if isGenRefresh {
		refresh = uuid.New().String()
	}

	return
}
