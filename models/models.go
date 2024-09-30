package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the user entity in the system.
type User struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name           string         `json:"name" gorm:"not null"`
	Email          string         `json:"email" gorm:"unique;not null"`
	HashedPassword string         `json:"-" gorm:"not null"` // Stored hashed password, not exposed in JSON
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hashedPassword)
	return nil
}

// CheckPassword verifies the provided password against the user's hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}

// UserRequest represents the request payload for creating or updating a user.
type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// UserResponse represents the response payload for user-related requests.
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginRequest represents the request payload for user login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Appointment represents the appointment entity in the system.
type Appointment struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title     string         `json:"title" gorm:"not null"`
	StartTime time.Time      `json:"start_time" gorm:"not null"`
	EndTime   time.Time      `json:"end_time" gorm:"not null"`
	Duration  time.Duration  `json:"duration" gorm:"not null"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	AppCode   string         `json:"App_code" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// AppointmentRequest represents the request payload for creating or updating an appointment.
type AppointmentRequest struct {
	Title     string        `json:"title" binding:"required"`
	StartTime time.Time     `json:"start_time" binding:"required"`
	EndTime   time.Time     `json:"end_time" binding:"required"`
	Duration  time.Duration `json:"duration" gorm:"not null"`
	UserID    uuid.UUID     `json:"user_id" binding:"required"`
}

// AppointmentResponse represents the response payload for appointment-related requests.
type AppointmentResponse struct {
	ID        uuid.UUID     `json:"id"`
	Title     string        `json:"title"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	UserID    uuid.UUID     `json:"user_id"`
	Duration  time.Duration `json:"duration" gorm:"not null"`
	AppCode   string        `json:"App_code" gorm:"not null"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// Booking represents a booking for an appointment.
type Booking struct {
	ID            uuid.UUID      `json:"id" gorm:"unique;type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        uuid.UUID      `json:"user_id" gorm:"type:uuid;not null"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	AppointmentID uuid.UUID      `json:"appointment_id" gorm:"type:uuid;not null"`
	Appointment   Appointment    `json:"appointment" gorm:"foreignKey:AppointmentID"`
	StartTime     time.Time      `json:"start_time" gorm:"not null"`
	EndTime       time.Time      `json:"end_time" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// BookingRequest represents the request payload for creating or updating a booking.
type BookingRequest struct {
	UserID        uuid.UUID `json:"user_id" binding:"required"`
	AppointmentID uuid.UUID `json:"appointment_id" binding:"required"`
	StartTime     time.Time `json:"start_time" binding:"required"`
	EndTime       time.Time `json:"end_time" binding:"required"`
}

// BookingResponse represents the response payload for booking-related requests.
type BookingResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	AppointmentID uuid.UUID `json:"appointment_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
