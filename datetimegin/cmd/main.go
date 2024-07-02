package main

import (
	"fmt"

	date_time "github.com/codescalersinternships/datetime-server-nabil/datetimegin/api/handler"
	"github.com/fvbock/endless"
)

func main() {
	r := date_time.SetupRouter()
	// Listen and Server in 0.0.0.0:8090
	r.Run(":8090")
	endless.ListenAndServe(":8090", r)

	fmt.Println("\nGracefully shutting down HTTP server...")
	fmt.Println("Shutdown complete.")
}
