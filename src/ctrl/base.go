package ctrl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error contains the error message
type Error struct {
	Error string `json:"error" example:"failed with error"`
}

// Success contains the success message
type Success struct {
	Success string `json:"success" example:"done with success"`
}

func renderErrorMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusBadRequest, &Error{Error: message})

	fmt.Println("API Error:", message)
	c.Abort()
}

func renderUnauthorizedError(c *gin.Context) {
	c.IndentedJSON(http.StatusUnauthorized, &Error{Error: "user is not authorized"})
	c.Abort()
}

func renderSuccessMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusOK, &Success{Success: message})
}

func renderError(c *gin.Context, err error) {
	renderErrorMessage(c, err.Error())
	c.Abort()
}

func renderJSON(c *gin.Context, obj interface{}) {
	c.IndentedJSON(http.StatusOK, obj)
}

func renderString(c *gin.Context, format string, values ...interface{}) {
	c.String(http.StatusOK, format, values...)
}
