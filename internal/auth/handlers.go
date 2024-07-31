package auth

import (
	"encoding/json"
	"inicio/internal/models"
	"net/http"
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
		return
	}

	token, err := Authenticate(credentials, ipAddress)
	if err != nil {
		http.Error(w, "Authentication failed: invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	err := deleteSession(user.Token)
	if err != nil {
		http.Error(w, "session could not be deleted", http.StatusInternalServerError)
		return
	}
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

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	err := deleteSession(user.Token)
	if err != nil {
		http.Error(w, "failed to delete session activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("session deleted successfully"))
	if err != nil {
		return
	}
}
