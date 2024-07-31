package auth

import (
	"database/sql"
	"errors"
	"inicio/internal/models"
	"inicio/internal/models/users"
	"time"
)

const sessionTimeout = 30 * time.Minute

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

func UpdateSession(token string) error {
	_, err := database.Exec("UPDATE sessions SET last_activity = ? WHERE token = ?", time.Now().Unix(), token)
	if err != nil {
		return err
	}
	return nil
}

func CheckSession(token string) (models.User, error) {
	session, err := getSessionByToken(token)
	if err != nil {
		return models.User{}, err
	}

	lastActivityTime := time.Unix(session.LastActivity, 0)
	if time.Since(lastActivityTime) > sessionTimeout {
		err = deleteSession(token)
		if err != nil {
			return models.User{}, err
		}
		return models.User{}, errors.New("session expired due to inactivity")
	}

	user, err := users.GetById(session.UserId)
	if err != nil {
		return models.User{}, err
	}

	user.Token = token
	return user, nil
}

func getSessionByToken(token string) (models.Session, error) {
	var session models.Session

	query := "SELECT token, user_id, last_activity FROM sessions WHERE token = ?"
	err := database.QueryRow(query, token).Scan(&session.Token, &session.UserId, &session.LastActivity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Session{}, errors.New("session not found")
		}
		return models.Session{}, err
	}

	return session, nil
}
