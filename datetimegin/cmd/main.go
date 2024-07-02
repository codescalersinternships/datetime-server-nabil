package main

import (
	"fmt"
	"log"

	date_time "github.com/codescalersinternships/datetime-server-nabil/datetimegin/api/handler"
	"github.com/fvbock/endless"
)

func main() {
	r := date_time.SetupRouter()
	// Listen and Server in 0.0.0.0:8090
	err := r.Run(":8090")
	if err != nil {
		log.Fatal("can't run server correctly", err)
	}
	err = endless.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal("can't end server gracefull", err)
	}
	fmt.Println("\nGracefully shutting down HTTP server...")
	fmt.Println("Shutdown complete.")
}
