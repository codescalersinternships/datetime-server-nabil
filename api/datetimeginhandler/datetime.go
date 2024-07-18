package datetime

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.GET("/datetime", GetDate)

	return r
}

// GetDateAndTime takes gin context
// and respond the current date and time in the format
// "2006-01-02 15:04:05" to the response writer.
// If the request is valid, it responds with a 200 OK status and the current date and time.
// If the request method is not GET, it responds with a 405 Method Not Allowed status.
// If the request URL is not "/datetime", it responds with a 404 Not Found status.
func GetDate(c *gin.Context) {
	currentDateTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprint(c.Writer, currentDateTime)
}
