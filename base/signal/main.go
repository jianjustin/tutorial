package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Create a new server
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	log.Println("Server is ready to handle requests at :8080")

	// Block until we receive an interrupt signal
	<-stop
	log.Println("Server is shutting down...")

	// Create a deadline to wait for the server to shut down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shut down the server: %v\n", err)
	}
	log.Println("Server stopped")
}
