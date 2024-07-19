package api

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Message string `json:"message"`
}

func SetupApiEndpoints() {
	http.HandleFunc("/api/data", sendData)
}

func sendData(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Message: "Hello from Go API",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
