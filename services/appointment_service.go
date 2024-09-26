package services

import (
	"github.com/google/uuid"
	"github.com/m13ha/appointment_master/db"
	models "github.com/m13ha/appointment_master/models"
)

// CreateAppointment creates a new appointment and saves it to the database.
func CreateAppointment(req models.AppointmentRequest) (*models.Appointment, error) {
	appointment := &models.Appointment{
		ID:     uuid.UUID{}, // Simulate UUID generation
		Title:  req.Title,
		UserID: req.UserID,
	}

	// Add appointment to DB
	if err := db.DB.Create(appointment).Error; err != nil {
		return nil, err
	}

	return appointment, nil
}

// GetUsersForAppointment retrieves users registered for a specific appointment.
func GetUsersForAppointment(appointmentID string) ([]models.User, error) {
	var users []models.User
	if err := db.DB.Where("appointment_id = ?", appointmentID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetCreatedAppointments retrieves all appointments created by the user.
func GetCreatedAppointments(userID string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	if err := db.DB.Where("user_id = ?", userID).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}
