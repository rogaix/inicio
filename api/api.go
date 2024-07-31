package api

import (
	"inicio/internal/auth"
	"inicio/internal/middleware"
	"net/http"
)

func SetupApiEndpoints() {
	http.Handle("/api/login", http.HandlerFunc(auth.LoginHandler))
	http.Handle("/api/register", http.HandlerFunc(auth.RegisterHandler))

	http.Handle("/api/updateSession", middleware.AuthMiddleware(http.HandlerFunc(auth.UpdateSessionHandler)))
	http.Handle("/api/deleteSession", middleware.AuthMiddleware(http.HandlerFunc(auth.DeleteSessionHandler)))
	http.Handle("/api/logout", middleware.AuthMiddleware(http.HandlerFunc(auth.LogoutHandler)))
}
