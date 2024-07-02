package datetime

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/datetime", GetDate)

	return r
}
func GetDate(c *gin.Context) {
	if c.Request.URL.Path != "/datetime" {
		http.NotFound(c.Writer, c.Request)
	}

	if c.Request.Method != "GET" {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	currentDateTime := time.Now().Format(time.RFC3339)
	fmt.Fprint(c.Writer, currentDateTime)
}
