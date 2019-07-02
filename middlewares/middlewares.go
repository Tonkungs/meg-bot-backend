package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
