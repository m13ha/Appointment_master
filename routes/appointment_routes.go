package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	models "github.com/m13ha/appointment_master/models"
	services "github.com/m13ha/appointment_master/services"
)

// CreateAppointment handles creating a new appointment
func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from context
	userIDStr, ok := r.Context().Value(UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	var appointmentReq models.AppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&appointmentReq); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	// Add user ID to the appointment request
	appointmentReq.UserID = userID

	// Validate required fields
	var validationErrors []models.ValidationError
	if appointmentReq.Title == "" {
		validationErrors = append(validationErrors, models.ValidationError{Field: "title", Message: "Title is required"})
	}
	if appointmentReq.StartTime.IsZero() {
		validationErrors = append(validationErrors, models.ValidationError{Field: "start_time", Message: "Start time is required"})
	}
	if appointmentReq.EndTime.IsZero() {
		validationErrors = append(validationErrors, models.ValidationError{Field: "end_time", Message: "End time is required"})
	}

	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewValidationErrorResponse(validationErrors...))
		return
	}

	appointment, err := services.CreateAppointment(appointmentReq)
	if err != nil {
		switch err.Error() {
		case "end time cannot be before start time":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "overlapping appointment exists":
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Failed to create appointment", http.StatusInternalServerError)
		}
		return
	}

	response := models.AppointmentResponse{
		ID:        appointment.ID,
		Title:     appointment.Title,
		StartTime: appointment.StartTime,
		EndTime:   appointment.EndTime,
		UserID:    appointment.UserID,
		AppCode:   appointment.AppCode,
		CreatedAt: appointment.CreatedAt,
		UpdatedAt: appointment.UpdatedAt,
	}

	w.WriteHeader(http.StatusCreated)
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
