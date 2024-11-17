package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondWithJSON sends a JSON response with a given status code
func RespondWithJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}

// HandleError is a method that checks for an error and responds accordingly
func HandleError(c *gin.Context, err error, message string) {
	if err != nil {
		// Send a standardized error response
		RespondWithError(c, http.StatusInternalServerError, message)
	}
}
