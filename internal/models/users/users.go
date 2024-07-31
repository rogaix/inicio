package users

import (
	"inicio/internal/db"
	"inicio/internal/models"
)

var database = db.GetDB()

func GetById(userId int) (models.User, error) {
	var user models.User

	err := database.QueryRow(""+
		"SELECT id, name, password "+
		"FROM users "+
		"WHERE id = ?", userId).Scan(&user.ID, &user.Name, &user.Password)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetByMailAddress(mailAddress string) (*models.User, error) {
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
