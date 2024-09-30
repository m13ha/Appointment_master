package services

import (
	"fmt"

	"github.com/m13ha/appointment_master/db"
	models "github.com/m13ha/appointment_master/models"
	"github.com/m13ha/appointment_master/utils"
)

// CreateAppointment creates a new appointment and saves it to the database.
func CreateAppointment(req models.AppointmentRequest) (*models.Appointment, error) {
	// Validate time range
	if req.EndTime.Before(req.StartTime) {
		return nil, fmt.Errorf("end time cannot be before start time")
	}

	// Check for overlapping appointments
	var count int64
	err := db.DB.Model(&models.Appointment{}).
		Where("user_id = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?))",
			req.UserID, req.StartTime, req.StartTime, req.EndTime, req.EndTime, req.StartTime, req.EndTime).
		Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("failed to check for overlapping appointments: %w", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("overlapping appointment exists")
	}

	appointment := &models.Appointment{
		Title:     req.Title,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserID:    req.UserID,
		AppCode:   utils.GenerateAppCode(),
		Duration:  req.Duration,
	}

	if err := db.DB.Create(appointment).Error; err != nil {
		return nil, fmt.Errorf("failed to create appointment: %w", err)
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
