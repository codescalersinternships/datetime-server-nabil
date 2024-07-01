package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetDate(w http.ResponseWriter, req *http.Request) {
	currentDateTime := time.Now().Format(time.RFC3339)
	fmt.Fprintf(w, currentDateTime)
}
func DateServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
func main() {

	http.HandleFunc("/datetime", GetDate)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
