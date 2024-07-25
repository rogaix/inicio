package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"inicio/internal/models"
	"os"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(credentials Credentials) (string, error) {
	user, err := getUserByMailAddress(credentials.Email)
	if err != nil || !checkPassword(user.Password, credentials.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Register(user models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return saveUser(user)
}

func checkPassword(hashedPassword, password string) bool {
	compareHashedPassword, err := hashPassword(password)
	if (hashedPassword != compareHashedPassword) || (err != nil) {
		return false
	}
	return true
}

func GenerateToken(user *models.User) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY_TOKEN"))

	atClaims := jwt.MapClaims{}
	atClaims["mailAddress"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(jwtKey)
	return token, err
}

func hashPassword(password string) (string, error) {
	value := []byte(password)
	hash := sha256.New()

	_, err := hash.Write(value)
	if err != nil {
		return "", err
	}

	hashedPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedPassword), nil
}

func saveUser(user models.User) error {
	// Save user to the repository
	return nil
}
