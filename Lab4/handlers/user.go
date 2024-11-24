package handlers

import (
	"Lab4/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: 1, Email: "user1@example.com", Password: "hashedpassword1"},
		{ID: 2, Email: "user2@example.com", Password: "hashedpassword2"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := models.User{
		ID:    1,
		Email: "user1@example.com",
	}

	if id != "1" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
