package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	models "github.com/m13ha/appointment_master/models"
	services "github.com/m13ha/appointment_master/services"
)

// CreateAppointment handles creating a new appointment
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointmentReq models.AppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&appointmentReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	appointment, err := services.CreateAppointment(appointmentReq)
	if err != nil {
		http.Error(w, "Failed to create appointment", http.StatusInternalServerError)
		return
	}

	response := models.AppointmentResponse{
		ID:     appointment.ID,
		Title:  appointment.Title,
		UserID: appointment.UserID,
	}
	json.NewEncoder(w).Encode(response)
}

// GetUsersRegisteredForAppointment retrieves all users registered for a specific appointment
func GetUsersRegisteredForAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentID := chi.URLParam(r, "id")
	users, err := services.GetUsersForAppointment(appointmentID)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetMyCreatedAppointments shows all appointments created by the user
func GetMyCreatedAppointments(w http.ResponseWriter, r *http.Request) {
	userID := "some_user_id"
	appointments, err := services.GetCreatedAppointments(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve appointments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(appointments)
}
