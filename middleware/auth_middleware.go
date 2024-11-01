// middleware/auth_middleware.go
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Define your secret key (store this securely in a real application)
	secretKey = []byte("your_secret_key")
)

// AuthMiddleware checks the JWT token in the request header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Split the header to get the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract user information from the token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["id"].(float64) // Assuming you're storing user ID in the token
			// Store user ID in context
			ctx := context.WithValue(r.Context(), "userID", int(userID))
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
