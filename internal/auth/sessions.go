package auth

import (
	"inicio/internal/models"
	"time"
)

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

func deleteSession(token string) error {
	_, err := database.Exec("DELETE FROM sessions WHERE token = ?", token)
	if err != nil {
		return err
	}

	return nil
}
