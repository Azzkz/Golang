package router

import (
	"Lab4/handlers"
	"Lab4/middleware"
	"Lab4/monitoring"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(monitoring.TrackMetrics)
	r.Use(middleware.SecurityHeaders)

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	r.HandleFunc("/csrf-protected", handlers.CSRFProtectedHandler).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")

	return r
}
