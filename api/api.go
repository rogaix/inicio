package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"inicio/internal/db"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

func SetupApiEndpoints() {
	http.HandleFunc("/api/data", sendData)
}

func sendData(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()

	var user User
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
