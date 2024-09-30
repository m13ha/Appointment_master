package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/m13ha/appointment_master/db"
	routes "github.com/m13ha/appointment_master/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get PORT from .env
	port := os.Getenv("PORT")

	// Connect to the database
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Auth routes
	r.Post("/login", routes.Login)
	r.Post("/logout", routes.Logout)

	// User routes
	r.Post("/users", routes.CreateUser)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(routes.AuthMiddleware)

		// Appointment routes
		r.Post("/appointments", routes.CreateAppointment)
		r.Get("/appointments/{id}/users", routes.GetUsersRegisteredForAppointment)
		r.Get("/appointments/my", routes.GetMyCreatedAppointments)
		r.Get("/appointments/registered", routes.GetRegisteredAppointments)
	})

	log.Printf("Starting Server on PORT %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
