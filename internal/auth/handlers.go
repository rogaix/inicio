package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"inicio/internal/models"
	"net/http"
	"os"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ipAddress, err := getIpAddress(r)
	if err != nil {
		http.Error(w, "IP Address cannot be retrieved", http.StatusBadRequest)
	}

	token, err := Authenticate(credentials, ipAddress)
	if err != nil {
		errMsg := fmt.Errorf("authentication failed: %v", err)
		http.Error(w, errMsg.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := Register(user); err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if len(strings.Split(tokenString, " ")) == 2 {
		tokenString = strings.Split(tokenString, " ")[1]
	} else {
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY_TOKEN")), nil
	})

	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mailAddress, ok := claims["mailAddress"].(string)
		if !ok {
			http.Error(w, "Invalid mail address claim", http.StatusBadRequest)
			return
		}

		// Use mail address to look up the user in DB
		user, err := getUserByMailAddress(mailAddress)
		if err != nil {
			return
		}

		// Generate a new JWT for the user.
		token, err := GenerateToken(user)
		if err != nil {
			return
		}

		// Return the new JWT to the client in JSON format.
		err = json.NewEncoder(w).Encode(map[string]string{"token": token})
		if err != nil {
			return
		}
	}
	return
}
