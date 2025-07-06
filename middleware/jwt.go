package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/wisnu-bdd/be-go-with-auth-template/config"
	"github.com/wisnu-bdd/be-go-with-auth-template/models"
	"github.com/golang-jwt/jwt/v5"
)

// Define a context key type to avoid collisions
type contextKey string
const userEmailKey contextKey = "userEmail"

// JWT is middleware that validates a JWT token in the Authorization header.
func JWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Inject email into request context
		ctx := context.WithValue(r.Context(), userEmailKey, claims.Email)
		r = r.WithContext(ctx)

		// Token is valid; continue
		next.ServeHTTP(w, r)
	}
}

// GetUserEmail retrieves the email from context in handlers
func GetUserEmail(r *http.Request) string {
	if email, ok := r.Context().Value(userEmailKey).(string); ok {
		return email
	}
	return ""
}
