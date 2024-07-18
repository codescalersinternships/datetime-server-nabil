package datetime

import (
	"fmt"
	"net/http"
	"time"
)

// GetDateAndTime takes in a response writer and a request
// and respond the current date and time in the format
// "2006-01-02 15:04:05" to the response writer.
// If the request is valid, it responds with a 200 OK status and the current date and time.
// If the request method is not GET, it responds with a 405 Method Not Allowed status.
// If the request URL is not "/datetime", it responds with a 404 Not Found status.
func GetDate(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/datetime" {
		http.NotFound(w, req)
	}

	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	currentDateTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprint(w, currentDateTime)
}
