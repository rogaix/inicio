package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"inicio/internal/auth"
	"inicio/internal/db"
	"inicio/internal/models"
	"log"
	"net/http"
)

type Data struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

func SetupApiEndpoints() {
	http.HandleFunc("/api/data", sendData)
	http.HandleFunc("/api/login", auth.LoginHandler)
	http.HandleFunc("/api/register", auth.RegisterHandler)
	http.HandleFunc("/api/refreshToken", auth.RefreshTokenHandler)
}

func sendData(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()

	var user models.User
	err := database.QueryRow(""+
		"SELECT id, name "+
		"FROM users "+
		"WHERE id = ?", 1).Scan(&user.ID, &user.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			sendJSONError(w, "User not found", http.StatusNotFound)
			return
		}
		log.Println("Error querying database:", err)
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := Data{
		Message: "Hello from Go API v2",
		User:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error encoding response:", err)
		sendJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		log.Println("Error encoding error response:", err)
	}
}
