package models

// ErrorResponse defines error structure for API failures
type ErrorResponse struct {
	Error string `json:"error" example:"Record not found"`
}

// SuccessMessage defines a generic success message
type SuccessMessage struct {
	Success string `json:"success" example:"Record deleted"`
}

type ErrorResponseCleanURL struct {
	Error string `json:"error" example:"Invalid input"`
}
