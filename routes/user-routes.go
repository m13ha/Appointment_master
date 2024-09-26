package routes

import (
	"encoding/json"
	"net/http"

	models "github.com/m13ha/appointment_master/models"
	services "github.com/m13ha/appointment_master/services"
)

// CreateUser handles creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(userReq)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	json.NewEncoder(w).Encode(response)
}

// GetRegisteredAppointments shows appointments a user registered for
func GetRegisteredAppointments(w http.ResponseWriter, r *http.Request) {
	userID := "some_user_id"
	appointments, err := services.GetRegisteredAppointments(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve appointments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}
