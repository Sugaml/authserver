package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3/server"
)

var jwtSecret = []byte("supersecretkey")

type JWTClaims struct {
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	jwt.StandardClaims
}

func ValidateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
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
