package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`          // "success" or "error"
	Message string      `json:"message"`         // Human-readable message
	Data    interface{} `json:"data,omitempty"`  // Optional, for success responses
	Error   *APIError   `json:"error,omitempty"` // Optional, for error responses
}

type APIError struct {
	Code    string `json:"code"`    // Custom error code
	Details string `json:"details"` // Detailed error message
}

// Success formats a standardized success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error formats a standardized error response
func Error(c *gin.Context, statusCode int, code, message, details string) {
	c.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
		Error: &APIError{
			Code:    code,
			Details: details,
		},
	})
}
