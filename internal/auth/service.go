package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"inicio/internal/models"
	"inicio/internal/models/users"
	"net"
	"net/http"
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

func Authenticate(credentials Credentials, ipAddress string) (string, error) {
	user, err := users.GetByMailAddress(credentials.Email)
	if err != nil || !checkPassword(user.Password, credentials.Password) {
		return "", ErrInvalidCredentials
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	err = saveSession(user, token, ipAddress)
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
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
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

func getIpAddress(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", errors.New("invalid ip address")
	}

	return ip, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func saveUser(user models.User) error {
	return nil
}
