package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"inicio/internal/auth"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header is required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		user, err := auth.CheckSession(tokenString)
		if err != nil {
			_ = fmt.Errorf("session invalid or expired %v, status: %d", err, http.StatusUnauthorized)
			return
		}

		// Renew token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY_TOKEN")), nil
		})

		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				exp := int64(claims["exp"].(float64))
				if time.Until(time.Unix(exp, 0)) < 5*time.Minute {
					newToken, err := auth.GenerateToken(&user)
					if err == nil {
						user.Token = newToken
						w.Header().Set("Authorization", "Bearer "+newToken)

						err = auth.UpdateSessionToken(tokenString, newToken)
						if err != nil {
							return
						}
					}
				}
			}
		}

		err = auth.UpdateSessionActivity(tokenString)
		if err != nil {
			http.Error(w, "failed to update session activity", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
