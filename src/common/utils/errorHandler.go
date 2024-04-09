package utils

import (
	commonModels "ctint-conv/src/common/models"
	"encoding/json"
	"net/http"
)

// HandleError returns a JSON error response with the provided error message.
func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":      nil,
		"isSuccess": false,
		"error":     err.Error(),
	})
}

func HandleErrorWithStatus(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":      nil,
		"isSuccess": false,
		"error":     err.Error(),
	})
}

// NewAppError creates a new instance of AppError with the given message.
func NewAppError(message string) *commonModels.AppError {
	return &commonModels.AppError{
		Message: message,
	}
}
