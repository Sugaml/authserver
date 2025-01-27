package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtSecret = []byte("supersecretkey")

// JWTClaims defines the structure of JWT token claims
type JWTClaims struct {
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	jwt.StandardClaims
}

type DBClientStore struct {
	DB *gorm.DB
}

type Client struct {
	ID     string `gorm:"primaryKey"`
	Secret string
	Domain string
}

func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=root password=secret dbname=auth port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the Client model
	if err := db.AutoMigrate(&Client{}); err != nil {
		return nil, err
	}

	return db, nil
}

// GetByID retrieves a client by its ID from the database
func (s *DBClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	var client Client
	if err := s.DB.First(&client, "id = ?", id).Error; err != nil {
		log.Println("Error fetching client:", err)
		return nil, errors.ErrInvalidClient
	}
	return &models.Client{
		ID:     client.ID,
		Secret: client.Secret,
		Domain: client.Domain,
	}, nil
}

// CreateClient adds a new client to the database
func (s *DBClientStore) CreateClient(id, secret, domain string) error {
	client := Client{
		ID:     id,
		Secret: secret,
		Domain: domain,
	}
	return s.DB.Create(&client).Error
}

func main() {

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// Token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Client memory store
	dbClientStore := &DBClientStore{DB: db}
	manager.MapClientStorage(dbClientStore)

	// Set custom JWT generator
	manager.MapAccessGenerate(&JWTAccessGenerate{})

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if err := srv.HandleTokenRequest(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		clientID := uuid.New().String()[:8]
		clientSecret := uuid.New().String()[:8]
		clientDomain := "http://localhost:9094"
		if err := dbClientStore.CreateClient(clientID, clientSecret, clientDomain); err != nil {
			http.Error(w, "Failed to create client: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientID, "CLIENT_SECRET": clientSecret})
	})

	http.HandleFunc("/protected", validateToken(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, srv))

	log.Fatal(http.ListenAndServe(":9096", nil))
}

// JWTAccessGenerate generates JWT tokens
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

func validateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr = tokenStr[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		f.ServeHTTP(w, r)
	})
}
