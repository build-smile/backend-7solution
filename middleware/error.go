package middleware

import (
	"github.com/build-smile/backend-7solution/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware is a middleware function to handle errors
func ErrorHandlerMiddleware(c *gin.Context) {
	// Run the next handler
	c.Next()

	// Check if there are any errors in the Gin context
	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		handleError(c, err)
	}
}

// handleError function to handle custom and generic errors
func handleError(c *gin.Context, err error) {
	// Check if the error is of type *service.CustomError
	if customErr, ok := err.(*utils.CustomError); ok {
		// Return the custom error's status code and message
		c.JSON(customErr.Code, gin.H{"error": customErr.Message})
	} else {
		// Fallback to 500 Internal Server Error for unknown errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
