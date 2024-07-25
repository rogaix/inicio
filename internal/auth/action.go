package auth

import (
	"inicio/internal/db"
	"inicio/internal/models"
	"time"
)

var database = db.GetDB()

func getUserByMailAddress(mailAddress string) (*models.User, error) {
	var user models.User

	err := database.QueryRow(""+
		"SELECT id, name, password "+
		"FROM users "+
		"WHERE email = ?", mailAddress).Scan(&user.ID, &user.Name, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func saveSession(user *models.User, token string, ipAddress string) error {
	currentTime := time.Now().Unix()
	_, err := database.Exec(""+
		"INSERT INTO sessions (token, user_id, ip_address, last_activity)"+
		"VALUES (?, ?, ?, ?)", token, user.ID, ipAddress, currentTime)

	if err != nil {
		return err
	}

	return nil
}

func insertUser(user models.User) error {
	_, err := database.Exec("INSERT INTO users (name, email, password) VALUES (?, ?)", user.Name, user.Email, user.Password)
	return err
}
