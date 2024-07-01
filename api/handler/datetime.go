package datetime

import (
	"fmt"
	"net/http"
	"time"
)

func GetDate(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/datetime" {
		http.NotFound(w, req)
	}

	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	currentDateTime := time.Now().Format(time.RFC3339)
	fmt.Fprint(w, currentDateTime)
}
