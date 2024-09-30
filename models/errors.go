package models

// ValidationError represents an error response for validation issues
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrorResponse is the structure for validation error responses
type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

// DatabaseErrorResponse represents an error response for database issues
type DatabaseErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// NewValidationErrorResponse creates a new ValidationErrorResponse
func NewValidationErrorResponse(errors ...ValidationError) ValidationErrorResponse {
	return ValidationErrorResponse{Errors: errors}
}

// NewDatabaseErrorResponse creates a new DatabaseErrorResponse
func NewDatabaseErrorResponse(message, code string) DatabaseErrorResponse {
	return DatabaseErrorResponse{Message: message, Code: code}
}
