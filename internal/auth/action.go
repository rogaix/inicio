package auth

import (
	"inicio/internal/db"
	"inicio/internal/models"
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

func insertUser(user models.User) error {
	_, err := database.Exec("INSERT INTO users (name, email, password) VALUES (?, ?)", user.Name, user.Email, user.Password)
	return err
}
