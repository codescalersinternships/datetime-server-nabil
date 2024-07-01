package main

import (
	"log"
	"net/http"

	date_time "github.com/codescalersinternships/datetime-server-nabil/api/handler"
)

func main() {
	http.HandleFunc("/datetime", date_time.GetDate)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
