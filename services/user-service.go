package services

import (
	"github.com/google/uuid"
	"github.com/m13ha/appointment_master/db"
	"github.com/m13ha/appointment_master/models"
)

// CreateUser creates a new user and saves it to the database.
func CreateUser(req models.UserRequest) (*models.User, error) {
	user := &models.User{
		ID:    uuid.UUID{}, // Simulate UUID generation
		Name:  req.Name,
		Email: req.Email,
	}

	// Add user to DB
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetRegisteredAppointments retrieves appointments registered by a user.
func GetRegisteredAppointments(userID string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	if err := db.DB.Where("user_id = ?", userID).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}
