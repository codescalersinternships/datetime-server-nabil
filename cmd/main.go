package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	date_time "github.com/codescalersinternships/datetime-server-nabil/api/handler"
)

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	server := &http.Server{
		Addr: ":8090",
	}
	go func() {
		fmt.Println("app running on 8090")
		http.HandleFunc("/datetime", date_time.GetDate)
		log.Fatal(http.ListenAndServe(":8090", nil))

	}()

	<-ctx.Done()
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatal("HTTP server shutdown error: ", err)
	}

	fmt.Println("HTTP server stopped.")

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go startHTTPServer(ctx, &wg)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	fmt.Println("\nGracefully shutting down HTTP server...")

	cancel()

	wg.Wait()

	fmt.Println("Shutdown complete.")
}
