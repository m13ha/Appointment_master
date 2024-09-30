package services

import (
	"github.com/m13ha/appointment_master/db"
	"github.com/m13ha/appointment_master/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user and saves it to the database.
func CreateUser(userReq models.UserRequest) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:           userReq.Name,
		Email:          userReq.Email,
		HashedPassword: string(hashedPassword),
	}

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
